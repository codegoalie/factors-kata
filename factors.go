package main

func Factors(input int) (factors []int) {
	candidate := 2

	for input != 1 {
		if input%candidate == 0 {
			factors = append(factors, candidate)
			input = input / candidate
		} else {
			candidate += 1
		}
	}

	return factors
}
