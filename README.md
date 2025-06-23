ZeroScope — Advanced APK Static Analyzer
ZeroScope is a high-performance static analysis tool for Android APKs. Built with Go, it efficiently extracts metadata, scans for sensitive components, and generates professional reports in multiple formats — all with cyberpunk polish and smart automation.

🚀 What's New in v1.2
- ✅ API call classification by HTTP method (GET, POST, PUT, DELETE, OTHER)
- ✅ Smart filtering to ignore system logs and irrelevant strings
- ✅ Automatic extraction and analysis of embedded archives (.zip inside assets/)
- ✅ Saves unclassified strings into unclassified.txt for future analysis
- ✅ Displays a terminal summary after report generation with pause before exit
- ✅ Clean parsing of human-readable strings only (filters binary-like lines)

🛠️ Requirements
- Go 1.20+
- Optional: install gofpdf for PDF export
go get github.com/jung-kurt/gofpdf



📦 Folder Structure
zeroscope/
├── APKs/               → Drop your `.apk` files here
├── reports/            → Auto-generated reports (JSON/HTML/PDF)
├── tmp/                → Temporary unpacked contents
├── core/               → Analysis and reporting logic
├── utils/              → Unzip, archive extraction, helpers
├── rules/              → Plugin system (optional, coming soon)
├── cmd/                → Entry point (main.go)
└── README.md



⚙️ Usage
go run ./cmd


- Place one or more APKs in the APKs/ directory
- Run the tool — it'll unpack, analyze, and generate:
- report.html
- report.json
- unclassified.txt

📋 Report Preview
📋 Summary:
  - Permissions:      12
  - Dangerous:        5
  - Secrets Found:    3
  - API Calls (GET):  6
  - API Calls (POST): 4
  - API Calls (OTHER):2
  - Sensitive Paths:  2
  - Suspicious Files: 1
  - Unclassified:     19


📡 About
Made with 💻 by scz0x

⚠️ Disclaimer
ZeroScope is intended solely for educational, research, and ethical security analysis purposes. Use it only on APK files you own or have explicit authorization to analyze.