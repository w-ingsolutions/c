package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

// FileExists returns whether a file exists
func P() *gofpdf.Fpdf {
	return gofpdf.New("P", "", "", "/usr/home/marcetin/go/src/github.com/w-ingsolutions/c/pkg/font")
}
