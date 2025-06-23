package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"zeroscope/core"
)

func main() {
	apkFolder := "APKs"
	reportBase := "reports"
	tmpFolder := "tmp"

	files, err := os.ReadDir(apkFolder)
	if err != nil {
		fmt.Println("[✗] Failed to read APKs folder:", err)
		return
	}

	for _, entry := range files {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".apk") {
			continue
		}

		name := entry.Name()
		apkPath := filepath.Join(apkFolder, name)
		baseName := strings.TrimSuffix(name, ".apk")

		fmt.Println("🔍 Analyzing:", apkPath)

		apkTempDir := filepath.Join(tmpFolder, baseName)
		err := core.UnzipAPK(apkPath, apkTempDir)
		if err != nil {
			fmt.Println("[✗] Failed to unpack:", err)
			continue
		}

		report := core.AnalyzeAPK(apkTempDir)

		outputDir := filepath.Join(reportBase, baseName)
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Println("[✗] Failed to create report folder:", err)
			continue
		}

		core.GenerateJSONReport(outputDir, report)
		core.GenerateHTMLReport(outputDir, report)
		core.GeneratePDFReport(outputDir, report)

		err = os.RemoveAll(apkTempDir)
		if err == nil {
			fmt.Println("🧹 Temporary files removed:", apkTempDir)
		} else {
			fmt.Println("[!] Failed to remove temp folder:", err)
		}

		fmt.Println("✅ Report saved:", outputDir)
	}

	fmt.Println("🎯 All APKs processed.")
}