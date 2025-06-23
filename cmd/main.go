package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"zeroscope/core"
	"zeroscope/utils"
)

func main() {
	fmt.Println("🔍 ZeroScope", core.Version)

	apkFolder := "APKs"
	outputBase := "reports"
	tmp := "tmp"

	files, err := os.ReadDir(apkFolder)
	if err != nil {
		fmt.Println("✗ Failed to read APK folder:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".apk") {
			continue
		}

		name := file.Name()
		base := strings.TrimSuffix(name, ".apk")
		apkPath := filepath.Join(apkFolder, name)
		tmpDir := filepath.Join(tmp, base)
		outDir := filepath.Join(outputBase, base)

		utils.UnzipAPK(apkPath, tmpDir)
		utils.ExtractArchives(tmpDir)

		report := core.AnalyzeAPK(tmpDir)
		report.Version = core.Version
		core.AnalyzeExtractedArchives(tmpDir, &report)

		os.MkdirAll(outDir, 0755)
		core.GenerateJSONReport(outDir, report)
		core.GenerateHTMLReport(outDir, report)
		core.SaveUnclassifiedText(report.Unclassified, outDir)

		fmt.Println("📋 Summary:")
		fmt.Printf("  - Permissions:      %d\n", len(report.Permissions))
		fmt.Printf("  - Dangerous:        %d\n", len(report.Dangerous))
		fmt.Printf("  - Secrets Found:    %d\n", len(report.SecretsFound))
		fmt.Printf("  - API Calls (GET):  %d\n", len(report.APICalls["GET"]))
		fmt.Printf("  - API Calls (POST): %d\n", len(report.APICalls["POST"]))
		fmt.Printf("  - API Calls (OTHER):%d\n", len(report.APICalls["OTHER"]))
		fmt.Printf("  - Sensitive Paths:  %d\n", len(report.SensitivePaths))
		fmt.Printf("  - Suspicious Files: %d\n", len(report.Suspicious))
		fmt.Printf("  - Unclassified:     %d\n", len(report.Unclassified))

		fmt.Print("\n🔸 Press ENTER to Exit..")
		fmt.Scanln()

		os.RemoveAll(tmpDir)
		fmt.Println("✅ Done:", base)
	}
}