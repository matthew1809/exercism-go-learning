package pangram

import (
	"unicode"
)

// IsPangram checks if the input string contains every letter of the alphabet at least once,
// and returns true or false based on wether the input matches this condition
func IsPangram(input string) bool {

	// Could generate this instead using the appropriate unicode number range
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	runeListInput := []rune(input)

	for _, char := range runeListInput {
		for i, v := range alphabet {
			if unicode.ToLower(char) == v {
				alphabet = append(alphabet[:i], alphabet[i+1:]...)
			}
		}
	}

	if len(alphabet) == 0 {
		return true
	}
	return false
}
