package proverb

import "fmt"

func generate_main_rhyme_sentance(input1 string, input2 string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", input1, input2)
}

func generate_end_rhyme_sentance(input string) string {
	return fmt.Sprintf("And all for the want of a %s.", input)
}

func Proverb(listToRhyme []string) []string {

	output := []string{}

	if len(listToRhyme) < 1 {
		return output
	} else if len(listToRhyme) == 1 {
		output = append(output, generate_end_rhyme_sentance(listToRhyme[0]))
	} else if len(listToRhyme) == 2 {
		output = append(output, generate_main_rhyme_sentance(listToRhyme[0], listToRhyme[1]))
		output = append(output, generate_end_rhyme_sentance(listToRhyme[0]))
	} else if len(listToRhyme) > 2 {
		for i := 0; i < len(listToRhyme)-1; i++ {
			output = append(output, generate_main_rhyme_sentance(listToRhyme[i], listToRhyme[i+1]))
		}
		output = append(output, generate_end_rhyme_sentance(listToRhyme[0]))
	}

	return output
}
