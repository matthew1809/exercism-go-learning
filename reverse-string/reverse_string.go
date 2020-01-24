package reverse

// Reverse takes a string and returns the exact reverse of that string
func Reverse(input string) string {
	chars := []rune(input)

	// take the first and the end index numbers equal to i and j respectively
	// keep looping while i is less than the end index number i.e. go till the end of the list of characters
	// each loop, increment the starting index number by one, reduce the end index number by one
	// each loop, the value of the character at the starting index number will whatever the value at the end index number is
	// each loop, the value of the character at the ending number will be whatever the value at the value of the character at the starting index is

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
