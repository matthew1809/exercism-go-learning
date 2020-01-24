package wordcount

import (
	"fmt"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	output := Frequency{}

	fmt.Println(phrase)
	var phraseArray []string

	if strings.Contains(phrase, ",") {
		fmt.Println("splitting on comma")
		phraseArray = strings.Split(phrase, ",")
		phraseArray = strings.Split(phrase, " ")
	} else if strings.Contains(phrase, " ") {
		fmt.Println("splitting on space")
		phraseArray = strings.Split(phrase, " ")
	} else {
		phraseArray = []string{phrase}
	}

	fmt.Println(phraseArray)

	// loop over each word in the phrase array
	for _, input_word := range phraseArray {

		// we haven't found the phrase in output yet
		found := false

		// loop over each word in output if there are any
		for i, _ := range output {

			// if we find the word already in output
			if input_word == i {
				found = true
				// increment the count for that word in output by one
				output[i] = output[i] + 1
			}
		}

		// if we reach here and found is still false then the word isn't in the output list
		if found == false {
			output[input_word] = 1
		}
	}

	return output
}
