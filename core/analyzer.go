package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AnalyzeAPK(apkFolder string) Report {
	report := Report{}

	manifestPath := filepath.Join(apkFolder, "AndroidManifest.xml")
	if _, err := os.Stat(manifestPath); err == nil {
		report.Permissions = ExtractPermissions(manifestPath)
		report.Dangerous = FilterDangerousPermissions(report.Permissions)
		report.DangerousSet = make(map[string]bool)
		for _, d := range report.Dangerous {
			report.DangerousSet[d] = true
		}
	}

	report.FileCounts, report.SizesMB = CountFileTypes(apkFolder)
	report.Suspicious = DetectSuspiciousFiles(apkFolder)

	dexStrings := ExtractDexStrings(apkFolder)

	report.SecretsFound = ScanSensitiveStringsFromList(dexStrings)
	report.APICalls = ScanForAPICallsGrouped(dexStrings)
	report.SensitivePaths = ScanForSensitivePaths(dexStrings)

	var classified []string
	classified = append(classified, report.SecretsFound...)
	for _, group := range report.APICalls {
		classified = append(classified, group...)
	}
	classified = append(classified, report.SensitivePaths...)

	seen := make(map[string]bool)
	for _, c := range classified {
		seen[c] = true
	}
	for _, line := range dexStrings {
		line = strings.TrimSpace(line)
		if line == "" || seen[line] {
			continue
		}
		report.Unclassified = append(report.Unclassified, line)
	}

	return report
}

func AnalyzeExtractedArchives(folder string, report *Report) {
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		lower := strings.ToLower(path)
		if strings.HasSuffix(lower, ".json") || strings.HasSuffix(lower, ".txt") || strings.HasSuffix(lower, ".xml") || strings.HasSuffix(lower, ".cfg") {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			lines := strings.Split(string(data), "\n")
			report.SecretsFound = append(report.SecretsFound, ScanSensitiveStringsFromList(lines)...)
			apiCalls := ScanForAPICallsGrouped(lines)
			for method, group := range apiCalls {
				report.APICalls[method] = append(report.APICalls[method], group...)
			}
			report.SensitivePaths = append(report.SensitivePaths, ScanForSensitivePaths(lines)...)
		}
		return nil
	})
}

func SaveUnclassifiedText(lines []string, outputPath string) {
	if len(lines) == 0 {
		return
	}
	txtPath := filepath.Join(outputPath, "unclassified.txt")
	f, err := os.Create(txtPath)
	if err != nil {
		fmt.Println("[✗] Failed to save unclassified.txt:", err)
		return
	}
	defer f.Close()
	for _, line := range lines {
		_, _ = f.WriteString(line + "\n")
	}
	fmt.Println("📎 Saved:", txtPath)
}