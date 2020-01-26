package isbn

import (
	"strings"
	"unicode"
)

func runesToInts(input []rune) []int {
	output := []int{}

	for _, v := range input {
		if unicode.IsNumber(v) {
			output = append(output, int(v-'0'))
		} else {
			if v == 88 {
				output = append(output, 10)
			}

		}
	}

	return output
}

func allDigitsOrXInLastPosition(input string) bool {
	for i, v := range []rune(input) {
		if i < len(input)-1 {
			if !unicode.IsNumber(v) {
				return false
			}
		} else if i == len(input)-1 {
			if !unicode.IsNumber(v) && v != 'X' {
				return false
			}
		}
	}

	return true
}

func IsValidISBN(input string) bool {
	t := strings.Replace(input, "-", "", -1)

	if len(t) <= 9 || len(t) > 10 || !allDigitsOrXInLastPosition(t) {
		return false
	}

	ints := runesToInts([]rune(t))

	if (ints[0]*10+ints[1]*9+ints[2]*8+ints[3]*7+ints[4]*6+ints[5]*5+ints[6]*4+ints[7]*3+ints[8]*2+ints[9]*1)%11 != 0 {
		return false
	}

	return true
}
