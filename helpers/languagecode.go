package helpers

import (
	"github.com/cufee/am-wg-proxy-next/types"
)

func GetLanguageCode(lang string) string {
	return types.GetLocale(lang)
}
