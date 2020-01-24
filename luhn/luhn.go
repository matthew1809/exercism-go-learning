package luhn

import (
	"strconv"
	"unicode"
)

// allDigitsAndSpaces, given a slice of runes, returns true if there are no runes which are not digits or spaces
func allDigitsAndSpaces(input []rune) bool {
	for _, v := range input {
		if !unicode.IsDigit(v) && !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

// Valid returns true if the input string without spaces is valid according to the Luhn algorithm, otherwise it returns false
func Valid(input string) bool {
	inputAsRunes := []rune(input)
	var inputAsInts []int

	// only accept strings with digits and spaces
	if !allDigitsAndSpaces(inputAsRunes) {
		return false
	}

	// convert to ints and strip spaces
	for _, v := range inputAsRunes {
		if !unicode.IsSpace(v) {
			intVal, err := strconv.Atoi(string(v))

			if err == nil {
				inputAsInts = append(inputAsInts, intVal)
			}
		}
	}

	// don't accept less than two ints
	if len(inputAsInts) < 2 {
		return false
	}

	// double every second number starting from the right end of the list
	for i := len(inputAsInts) - 2; i >= 0; i = i - 2 {
		var result int
		double := inputAsInts[i] + inputAsInts[i]

		if double > 9 {
			result = double - 9
		} else {
			result = double
		}
		inputAsInts[i] = result
	}

	var sum int

	// compute sum of all ints in list
	for _, v := range inputAsInts {
		sum = sum + v

	}

	// if sum divides evenly into ten, list is valid
	if sum%10 == 0 {
		return true
	}

	// if reaching this point, list is not valid
	return false

}
