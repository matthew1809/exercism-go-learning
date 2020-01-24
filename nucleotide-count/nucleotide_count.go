package dna

import (
	"errors"
	"strings"
)

type Histogram map[rune]int

type DNA string

func IsValidNuclide(s string) bool {
	for _, r := range s {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return false
		}
	}
	return true
}

func (d DNA) Counts() (Histogram, error) {
	if len(d) > 0 && !IsValidNuclide(string(d)) {
		return nil, errors.New("invalid nuclides found")
	}

	h := make(Histogram)
	h['A'] = strings.Count(string(d), "A")
	h['C'] = strings.Count(string(d), "C")
	h['G'] = strings.Count(string(d), "G")
	h['T'] = strings.Count(string(d), "T")

	return h, nil
}
