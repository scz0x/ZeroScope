package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var sensitiveKeywords = []string{
	"apikey=", "client_secret", "Authorization:", "Bearer", "token=", "access_key", "secret=", "auth_token",
}

func ScanSensitiveStrings(baseDir string) []string {
	var findings []string

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		// تحديد أنواع الملفات القابلة للفحص
		switch strings.ToLower(filepath.Ext(path)) {
		case ".dex", ".xml", ".json", ".txt", ".js", ".conf", ".cfg", ".ini", ".so":
			// قابل للفحص
		default:
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			for _, key := range sensitiveKeywords {
				if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
					entry := fmt.Sprintf("%s (line %d): %s", path, lineNum, strings.TrimSpace(line))
					fmt.Printf("\x1b[33m[!] Found suspicious string:\x1b[0m %s\n", entry)
					findings = append(findings, entry)
				}
			}
			lineNum++ // <-- هذا السطر مهم لزيادة رقم السطر
		}
		return nil
	})

	return findings
}