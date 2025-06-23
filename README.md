# 🧠 ZeroScope — Advanced APK Static Analyzer

ZeroScope is a high-performance static analysis tool for Android APKs. Built with Go, it efficiently extracts metadata, scans for sensitive components, and generates professional reports in multiple formats (JSON, HTML, PDF) — all with cyberpunk polish and smart automation.

---

## 🚀 Features

- 📦 **Unpacks and analyzes APK files** (reads `.dex`, `.so`, `AndroidManifest.xml`, assets, etc.)
- 🔐 **Extracts and flags dangerous permissions** from binary AndroidManifest (AXML format)
- 🧬 **Scans for secrets** like API keys, tokens, and sensitive strings
- ⚠️ **Detects suspicious `.so` libraries** often linked to exploits or injections
- 📊 **Generates clean reports** in:
  - `report.json` — machine-readable
  - `report.html` — dark-mode with branding support
  - `report.pdf` — ready-to-print & share
- 🔧 Easily extensible: drop APKs into `/APKs`, get reports in `/reports`

---

## 🛠️ Requirements

- Go 1.20+
- Optional: install `gofpdf` for PDF export

```bash
go get github.com/jung-kurt/gofpdf
```

---

## 📁 Folder Structure

```
zeroscope/
├── APKs/               → Drop your `.apk` files here
├── reports/            → Auto-generated reports (JSON/HTML/PDF)
├── core/               → Analysis & reporting logic
│   └── report.go
├── utils/              → Helpers (e.g., unzip, string scanning)
├── cmd/
│   └── main.go         → Main entry point
```

---

## ⚙️ Usage

```bash
go run ./cmd
```

1. Add one or more `.apk` files to `APKs/`
2. Tool unpacks each, analyzes contents, generates report per file
3. Reports saved to `reports/<apk_name>/report.{json,html,pdf}`

---

## 📄 Sample Output

- `report.json`: raw structured data
- `report.html`: formatted with tables and dark theme
- `report.pdf`: clean printable output 

---

## 📡 About

Made with 💻 by [scz0x](https://t.me/SCZ0X_CH)  

---

## ⚠️ Disclaimer
**Disclaimer:**  
ZeroScope is intended solely for educational, research, and ethical security analysis purposes. Use it only on APK files you own or have explicit authorization to analyze. The authors are not responsible for misuse or any consequences arising from unauthorized use.
