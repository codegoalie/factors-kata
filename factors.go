package factors

func Generate(n int) (factors []int) {
	candidate := 2
	remaining := n
	for remaining > 1 {
		if remaining%candidate == 0 {
			factors = append(factors, candidate)
			remaining /= candidate
		} else {
			candidate += 1
		}
	}

	return factors
}
