package core

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func GenerateHTMLFromTemplate(outputPath string, report Report) {
	tmplPath := "core/templates/report_template.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("[✗] Failed to load HTML template:", err)
		if Logger != nil {
			Logger.Println("Template Load Error:", err)
		}
		return
	}

	outFile := filepath.Join(outputPath, "report.html")
	f, err := os.Create(outFile)
	if err != nil {
		fmt.Println("[✗] Failed to create HTML report:", err)
		if Logger != nil {
			Logger.Println("HTML File Creation Error:", err)
		}
		return
	}
	defer f.Close()

	if err := tmpl.Execute(f, report); err != nil {
		fmt.Println("[✗] Failed to render HTML:", err)
		if Logger != nil {
			Logger.Println("HTML Render Error:", err)
		}
		return
	}

	fmt.Println("🌐 HTML report saved:", outFile)
	if Logger != nil {
		Logger.Println("HTML report successfully saved:", outFile)
	}
}