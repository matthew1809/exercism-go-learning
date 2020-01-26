package acronym

import (
	"strings"
	"unicode"
)

func containsHyphenOrUnderscore(input string) (bool, string) {
	asRunes := []rune(input)

	for _, v := range asRunes {
		if v == '_' {
			return true, "_"
		} else if v == '-' {
			return true, "-"
		}
	}
	return false, ""
}

func Abbreviate(s string) string {

	acr := ""
	list := strings.Split(s, " ")

	for _, v := range list {
		b, contains := containsHyphenOrUnderscore(v)

		if b {
			splitString := strings.Split(v, contains)

			for _, s := range splitString {
				if len(s) > 1 {
					runes := []rune(s)
					first := runes[0]
					acr = acr + string(unicode.ToUpper(first))
				}
			}
		} else {
			runes := []rune(v)
			first := runes[0]
			acr = acr + string(unicode.ToUpper(first))
		}
	}
	return acr
}
