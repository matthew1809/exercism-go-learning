package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for key, value := range input {
		for _, innerValue := range value {
			output[strings.ToLower(innerValue)] = key
		}
	}

	return output
}
