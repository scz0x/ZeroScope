package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateJSONReport(outputPath string, report Report) {
	outFile := filepath.Join(outputPath, "report.json")
	f, err := os.Create(outFile)
	if err != nil {
		fmt.Println("[✗] Failed to create JSON report:", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(report); err != nil {
		fmt.Println("[✗] Failed to write JSON:", err)
		return
	}

	fmt.Println("📝 JSON report saved:", outFile)
}

func GenerateHTMLReport(outputPath string, report Report) {
	GenerateHTMLFromTemplate(outputPath, report)
}