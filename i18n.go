package i18n

import (
	"os"
	"strings"
)

var (
	translations = map[string]Translation{}
	language     = "en_US"
)

type Translation map[string]string

func AddTranslation(translation Translation, language string) {
	translations[language] = translation
}

func LoadTranslator() {
	lang := os.Getenv("LANG")
	if lang != "" {
		language = lang[:strings.Index(lang, ".")]
	}
}

func SetLanguage(lang string) {
	language = lang
}

func Tr(key string) string {
	value, ok := translations[language][key]
	if ok {
		return value
	} else {
		return key
	}
}
