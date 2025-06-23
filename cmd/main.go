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
	apksDir := "APKs"
	files, err := os.ReadDir(apksDir)
	if err != nil {
		fmt.Println("[✗] APKs folder not found:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".apk") {
			continue
		}

		apkPath := filepath.Join(apksDir, file.Name())
		fmt.Println("\n[*] Scanning", apkPath)

		tmpDir := "tmp"
		os.MkdirAll(tmpDir, 0755)

		if err := utils.UnzipAPK(apkPath, tmpDir); err != nil {
			fmt.Println("[✗] Failed to unzip:", err)
			continue
		}

		counts, sizes, so := core.AnalyzeExtractedFiles(tmpDir)

		manifest := filepath.Join(tmpDir, "AndroidManifest.xml")
		var permissions []string
		if _, err := os.Stat(manifest); err == nil {
			permissions = core.AnalyzePermissions(manifest)
		}

		secrets := core.ScanSensitiveStrings(tmpDir)

		report := core.Report{
			FileCounts:   counts,
			SizesMB:      sizes,
			Suspicious:   so,
			Permissions:  permissions,
			Dangerous:    nil,
			SecretsFound: secrets,
		}

		reportDir := filepath.Join("reports", strings.TrimSuffix(file.Name(), ".apk"))
		core.GenerateReportJSON(reportDir, report)
		core.GenerateHTMLReport(reportDir, report)
		core.GeneratePDFReport(reportDir, report)

		os.RemoveAll(tmpDir)
		fmt.Println("[✓] Finished", file.Name())
	}

	fmt.Println("\n[✓] Analysis complete.")
	fmt.Print("Press ENTER to exit...")
	fmt.Scanln()
}