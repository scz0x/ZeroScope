```markdown
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
- 🖼️ Custom branding with your logo (`logo.png`) and Telegram channel
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
- `report.pdf`: clean printable output with your logo and contact info

<p align="center"><img src="reports/example/logo.png" width="120" /></p>

---

## 📡 About

Made with 💻 by [scz0x](https://t.me/SCZ0X_CH)  
Tool reports are tagged with your brand and link back to your Telegram channel.

---

## ⚠️ Disclaimer

ZeroScope is built for educational and ethical analysis. Please use it responsibly, only on files you own or have explicit permission to inspect.
```

---