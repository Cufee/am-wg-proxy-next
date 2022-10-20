package helpers

import (
	"strings"

	"github.com/byvko-dev/am-types/wargaming/v2/glossary"
)

func GetLanguageCode(lang string) string {
	for _, code := range glossary.AllLanguages {
		if code == strings.ToLower(lang) {
			return lang
		}
	}
	return glossary.LangEN
}
