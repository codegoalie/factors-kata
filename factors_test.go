package main

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
