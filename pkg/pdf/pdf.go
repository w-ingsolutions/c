package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

// FileExists returns whether a file exists
func P(filePath string) *gofpdf.Fpdf {
	return gofpdf.New("P", "", "", "./../font")
}
