package translate

import (
	"github.com/bregydoc/gtranslate"
)

type Translation struct {
	Source, Result string
}

func (t *Translation) T(text string) string {
	if t.Source != t.Result {
		translated, err := gtranslate.TranslateWithParams(
			text,
			gtranslate.TranslationParams{
				From: t.Source,
				To:   t.Result,
			},
		)
		if err != nil {
			translated = err.Error()
		}
		text = translated
	}
	return text
}
