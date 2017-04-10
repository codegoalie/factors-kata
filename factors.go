package factors

// PrimeFactors returns a slice of integers which are the prime factors of the
// given integer
func PrimeFactors(x int) []int {
	factors := []int{}

	candidate := 2
	for candidate <= x {
		if x%candidate == 0 {
			factors = append(factors, candidate)
			x = x / candidate
		} else {
			candidate++
		}
	}

	return factors
}
