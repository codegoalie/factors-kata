package factors

import (
	"bytes"
	"encoding/binary"
)

// Fuzz validates the provided imput in order for it to be provided to
// PrimeFactors
func Fuzz(data []byte) int {
	x, err := binary.ReadVarint(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	a := PrimeFactors(int(x))
	b := PrimeFactors(int(x))
	for i, aElement := range a {
		if aElement != b[i] {
			panic("a != b")
		}
	}

	return 1
}
