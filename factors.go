package factors

func Generate(n int) (factors []int) {
	candidate := 2
	for n != 1 {
		if n%candidate == 0 {
			n /= candidate
			factors = append(factors, candidate)
		} else {
			candidate += 1
		}
	}
	return factors
}
