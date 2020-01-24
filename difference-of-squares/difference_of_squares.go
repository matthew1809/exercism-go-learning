package diffsquares

import "fmt"

// SquareOfSum takes an integer and returns the square of the sum of all the numbers up to and including the integer
func SquareOfSum(count int) int {
	fmt.Println("input", count)

	sum := count * ((1 + count) / 2)
	fmt.Println("result of sum", sum)

	squareOfSum := sum * sum
	fmt.Println("square of sum", squareOfSum)
	return int(squareOfSum)
}

// SumOfSquares takes an integer and returns the sum of the squares of each of the numbers upto and including the integer
func SumOfSquares(count int) int {
	var sum int
	i := 0
	for i <= count {
		sum = sum + (i * i)
		i++
	}
	return sum
}

// Difference takes an integer and returns the SquareOfSum minus SumOfSquares for the given integer
func Difference(count int) int {
	return SquareOfSum(count) - SumOfSquares(count)
}
