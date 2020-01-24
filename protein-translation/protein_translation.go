package protein

import "errors"

var ErrStop error = errors.New("stop found")

var ErrInvalidBase error = errors.New("invalid base")

var lookUpTable = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
	"STOP":          {"UAA", "UAG", "UGA"},
}

func FromCodon(input string) (string, error) {
	var output string
	var found bool

	result := codonToProtein(input)

	if result != "STOP" && result != "" {
		output = codonToProtein(input)
		found = true
	} else if result == "STOP" {
		return "", ErrStop
	}

	if found == true {
		return output, nil
	} else {
		return "", ErrInvalidBase
	}
}

func codonToProtein(codon string) string {
	for key, val := range lookUpTable {
		for _, string := range val {
			if codon == string {
				return key
			}
		}
	}

	return ""
}

func FromRNA(input string) ([]string, error) {
	var output []string
	numOfStrings := len(input) / 3
	counter := 0

	var error error = nil

	for i := 0; i < numOfStrings; i++ {
		if counter <= (len(input) - 3) {
			codon := input[counter : counter+3]
			counter = counter + 3

			result := codonToProtein(codon)

			if result == "STOP" {
				return output, nil
			} else if result == "" {
				error = ErrInvalidBase
			} else {
				output = append(output, result)
			}
		}
	}

	if error == nil {
		return output, nil
	} else {
		return output, error
	}
}
