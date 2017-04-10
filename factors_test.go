package factors

import (
	"testing"
)

func TestFactors(t *testing.T) {
	cases := map[int][]int{
		1:  []int{},
		2:  []int{2},
		3:  []int{3},
		4:  []int{2, 2},
		5:  []int{5},
		6:  []int{2, 3},
		10: []int{2, 5},
	}

	for toFactor, expected := range cases {
		actual := PrimeFactors(toFactor)
		if len(actual) != len(expected) {
			t.Errorf("Cardinality of PrimeFactors(%d) == %d; want %d",
				toFactor,
				len(actual),
				len(expected))
		} else {
			for i, actualElement := range actual {
				if actualElement != expected[i] {
					t.Errorf("PrimeFactors(%d) == %q; want %q", toFactor, actual, expected)
				}
			}
		}
	}
}

func BenchmarkFactors(b *testing.B) {

	for i := 0; i < b.N; i++ {
		actual := PrimeFactors(10)
		expected := []int{2, 5}
		if len(actual) != len(expected) {
			b.Errorf("Cardinality of PrimeFactors(%d) == %d; want %d",
				10,
				len(actual),
				len(expected))
		} else {
			for i, actualElement := range actual {
				if actualElement != expected[i] {
					b.Errorf("PrimeFactors(%d) == %q; want %q", 10, actual, expected)
				}
			}
		}
	}
}
