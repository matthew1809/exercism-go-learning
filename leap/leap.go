// Package leap helps to calculate whether a year is a leap year or not
package leap

// isEvenlyDivisibleBy checks the modulus of two given inputs and returns true or false based on whether the first parameter can be evenly divided by the second parameter
func isEvenlyDivisibleBy(numToCheck int, byNumber int) bool {
	return numToCheck%byNumber == 0
}

// IsLeapYear takes a number and returns true or false based on whether that year is a valid leap year in the gregorian calendar
func IsLeapYear(numToCheck int) bool {

	evenlyDividedBy4 := isEvenlyDivisibleBy(numToCheck, 4)

	if !evenlyDividedBy4 {
		return false
	}

	evenlyDividedBy100 := isEvenlyDivisibleBy(numToCheck, 100)
	evenlyDividedBy400 := isEvenlyDivisibleBy(numToCheck, 400)

	if evenlyDividedBy100 && !evenlyDividedBy400 {
		return false
	} else if evenlyDividedBy100 && evenlyDividedBy400 {
		return true
	} else {
		return true
	}

}
