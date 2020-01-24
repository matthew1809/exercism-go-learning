package accumulate

// Accumulate takes a list of strings and a function, it applies the function to each string, returning a list of the transformed strings
func Accumulate(input []string, functionToApply func(string) string) []string {
	var output []string

	for _, s := range input {
		output = append(output, functionToApply(s))
	}

	return output
}
