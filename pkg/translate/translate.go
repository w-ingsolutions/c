package translate

import (
	"github.com/bregydoc/gtranslate"
)

func translate(text, source, result string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: source,
			To:   result,
		},
	)
	return translated, err
}
