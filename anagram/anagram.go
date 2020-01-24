package anagram

import (
	"strings"
	"unicode"
)

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func contains(char rune, word string) (bool, int) {
	for i, v := range word {
		if unicode.ToLower(char) == unicode.ToLower(v) {
			return true, i
		}
	}
	return false, -1
}

func isWordItself(input string, candidate string) bool {
	if strings.ToLower(input) == strings.ToLower(candidate) {
		return true
	}

	return false
}

func Detect(input string, candidates []string) []string {
	validCandidates := []string{}

	// loop over the candidates
	for _, candidate := range candidates {

		if !isWordItself(input, candidate) {

			inputDuplicate := input

			found := []rune{}

			inputToUse := input

			// loop over each character in a candidate
			for _, char := range []rune(candidate) {

				// if the character is in the word, append it to a list
				wordContainsChar, index := contains(char, inputToUse)

				if wordContainsChar {
					// add the character to found list
					found = append(found, char)

					// remove the character
					inputToUse = inputToUse[:index] + inputToUse[index+1:]
				}
			}

			// if there are as many characters in found as there are in the word, it's an anagram
			if (len(found) == len(inputDuplicate)) && (len(found) == len(candidate)) {
				validCandidates = append(validCandidates, candidate)
			}
		}
	}

	return validCandidates
}
