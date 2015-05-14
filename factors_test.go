package main

import "testing"

func TestFactors(t *testing.T) {
	cases := map[int][]int{
		1:  []int{},
		2:  []int{2},
		3:  []int{3},
		4:  []int{2, 2},
		5:  []int{5},
		6:  []int{2, 3},
		7:  []int{7},
		8:  []int{2, 2, 2},
		9:  []int{3, 3},
		10: []int{2, 5},
		11: []int{11},
		12: []int{2, 2, 3},
		13: []int{13},
		14: []int{2, 7},
		15: []int{3, 5},
	}

	for number, expected := range cases {
		actual := Factors(number)
		if !compare(expected, actual) {
			t.Errorf("Factor(%v) == %v, want %v", number, actual, expected)
		}
	}
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, value := range a {
		if value != b[i] {
			return false
		}
	}
	return true
}
