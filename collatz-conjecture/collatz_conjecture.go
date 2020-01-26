package collatzconjecture

import "errors"

func isEven(input int) bool {
	if input%2 == 0 {
		return true
	}
	return false
}

func CollatzConjecture(input int) (int, error) {

	count := 0
	if input == 1 {
		return count, nil
	} else if input < 1 {
		return count, errors.New("can't work with numbers less than one")
	}

	for input > 1 {
		if isEven(input) {
			input = input / 2
		} else {
			input = (input * 3) + 1
		}
		count++
	}

	return count, nil
}
