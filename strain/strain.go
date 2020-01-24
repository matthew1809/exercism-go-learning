package strain

type Ints []int
type Lists [][]int
type Strings []string

func (input Ints) Keep(function func(int) bool) Ints {
	output := []int{}

	if len(input) == 0 {
		return nil
	}

	for _, num := range input {
		if function(num) {
			output = append(output, num)
		}
	}

	return output
}

func (input Ints) Discard(function func(int) bool) Ints {
	output := []int{}

	if len(input) == 0 {
		return nil
	}

	for _, num := range input {
		if !function(num) {
			output = append(output, num)
		}
	}

	return output
}

func (input Strings) Keep(function func(string) bool) Strings {
	output := []string{}

	for _, str := range input {
		if function(str) {
			output = append(output, str)
		}
	}

	return output
}

func (input Lists) Keep(function func([]int) bool) Lists {
	output := [][]int{}

	for _, list := range input {
		if function(list) {
			output = append(output, list)
		}
	}
	return output
}
