package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Report struct {
	FileCounts   map[string]int     `json:"file_counts"`
	SizesMB      map[string]float64 `json:"sizes_mb"`
	Suspicious   []string           `json:"suspicious_so"`
	Permissions  []string           `json:"permissions"`
	Dangerous    []string           `json:"dangerous_permissions"`
	SecretsFound []string           `json:"secrets_found"`
}

func GenerateReportJSON(outputDir string, counts map[string]int, sizes map[string]float64, suspicious []string, permissions []string, secrets []string) {
	var danger []string
	for _, p := range permissions {
		if isDangerous(p) {
			danger = append(danger, p)
		}
	}

	report := Report{
		FileCounts:   counts,
		SizesMB:      sizes,
		Suspicious:   suspicious,
		Permissions:  permissions,
		Dangerous:    danger,
		SecretsFound: secrets,
	}

	os.MkdirAll(outputDir, 0755)
	outPath := filepath.Join(outputDir, "report.json")
	file, err := os.Create(outPath)
	if err != nil {
		fmt.Println("[✗] Failed to write report:", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(report)
	if err != nil {
		fmt.Println("[✗] Failed to encode report:", err)
		return
	}

	fmt.Println("[✓] report.json generated at", outPath)
}

func isDangerous(p string) bool {
	dangerous := []string{
		"android.permission.READ_SMS",
		"android.permission.SYSTEM_ALERT_WINDOW",
		"android.permission.REQUEST_INSTALL_PACKAGES",
		"android.permission.CALL_PHONE",
		"android.permission.RECORD_AUDIO",
		"android.permission.READ_CONTACTS",
		"android.permission.WRITE_EXTERNAL_STORAGE",
		"android.permission.READ_PHONE_STATE",
	}
	for _, d := range dangerous {
		if strings.EqualFold(p, d) {
			return true
		}
	}
	return false
}