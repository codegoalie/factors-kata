package factors

import "testing"

func TestGenerate(t *testing.T) {
	cases := make(map[int][]int)
	cases[1] = []int{}
	cases[2] = []int{2}
	cases[3] = []int{3}
	cases[4] = []int{2, 2}
	cases[5] = []int{5}
	cases[6] = []int{2, 3}
	cases[7] = []int{7}
	cases[8] = []int{2, 2, 2}
	cases[9] = []int{3, 3}

	for input, expected := range cases {
		got := Generate(input)
		if !sliceEquals(expected, got) {
			t.Errorf("Generate(%d) == %v, want %v", input, got, expected)
		}
	}
}

func sliceEquals(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i, lhs := range x {
		if lhs != y[i] {
			return false
		}
	}

	return true
}
