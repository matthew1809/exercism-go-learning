// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

type Kind string

const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// isValidTriangle returns true or false based on whether three points can ake a valid triangle
func isValidTriangle(a, b, c float64) bool {

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if ((a + b) < c) || ((b + c) < a) || ((a + c) < b) {
		return false
	}

	return true
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isValidTriangle(a, b, c) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else if a != b && b != c {
		k = Sca
	}

	return k
}
