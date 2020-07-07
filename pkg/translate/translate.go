package translate

import (
	"github.com/bregydoc/gtranslate"
)

type Translation struct {
	source, result string
}

func (t *Translation) translate(text string) string {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: t.source,
			To:   t.result,
		},
	)
	if err != nil {
		translated = err.String()
	}
	return translated
}
