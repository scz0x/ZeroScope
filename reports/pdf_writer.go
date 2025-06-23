package report

import "github.com/jung-kurt/gofpdf"

func GeneratePDF(path string, permissions []string, secrets []string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "ZeroScope Analysis Report")
    pdf.Ln(12)

    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, "Permissions Found:")
    pdf.Ln(8)
    for _, p := range permissions {
        pdf.Cell(10, 10, "- "+p)
        pdf.Ln(6)
    }

    pdf.Ln(6)
    pdf.Cell(40, 10, "Secrets Found:")
    pdf.Ln(8)
    for _, s := range secrets {
        pdf.MultiCell(0, 6, "- "+s, "", "", false)
    }

    return pdf.OutputFileAndClose(path)
}
