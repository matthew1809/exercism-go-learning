package romannumerals

import "fmt"

import "errors"

func CountDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func ToRomanNumeral(input int) (output string, err error) {

	// I = 1
	// V = 5
	// X = 10
	// L = 50
	// C = 100
	// D = 500
	// M = 1000

	fmt.Println((input))

	var outputString string

	if input <= 0 {
		return "", errors.New("Cannot contain a number less than or equal to zero")
	}

	if input < 5 {
		for i := 0; i < input; i++ {
			outputString = outputString + "I"
		}
		return outputString, nil
	} else {
		return "ABC", nil
	}
}
