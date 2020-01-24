package listops

type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

func (input IntList) Foldr(function binFunc, inputInt int) int {
	output := 0

	if len(input) > 0 {
		input[0] = function(inputInt, input[0])
		output = input[0]
		for i := 0; i < len(input)-1; i++ {
			output = function(input[i+1], output)
		}
	} else {
		output = inputInt
		return output
	}
	return output
}

func (input IntList) Foldl(function binFunc, inputInt int) int {
	output := 0

	if len(input) > 0 {
		input[0] = function(inputInt, input[0])
		output = input[0]
		for i := 0; i < len(input)-1; i++ {
			output = function(output, input[i+1])
		}
	} else {
		output = inputInt
		return output
	}
	return output
}

func (input IntList) Filter(function predFunc) IntList {
	output := IntList{}

	for _, num := range input {
		if function(num) {
			output = append(output, num)
		}
	}

	return output
}

func (input IntList) Map(function unaryFunc) IntList {
	output := IntList{}

	if len(input) == 0 {
		return input
	}

	for _, num := range input {
		output = append(output, function(num))
	}

	return output
}

func (input IntList) Reverse() IntList {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}

	return input
}

func (input IntList) Append(toAppend IntList) IntList {
	output := IntList{}
	output = append(input, toAppend...)
	return output
}

func (input IntList) Concat(toAppend []IntList) IntList {
	for _, list := range toAppend {
		input = append(input, list...)
	}

	return input
}

func (input IntList) Length() int {
	return len(input)
}
