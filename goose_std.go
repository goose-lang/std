package std

import (
	"github.com/tchajed/goose/machine"
)

// Test if the two byte slices are equal.
func BytesEqual(x []byte, y []byte) bool {
	xlen := len(x)
	if xlen != len(y) {
		return false
	}
	var i = uint64(0)
	var retval = true
	for i < uint64(xlen) {
		if x[i] != y[i] {
			retval = false
			break
		}
		i += 1
		continue
	}
	return retval
}

// Compute the sum of two numbers, `Assume`ing that this does not overflow.
// *Use with care*, assumptions are trusted and should be justified!
func SumAssumeNoOverflow(x uint64, y uint64) uint64 {
	machine.Assume(x+y >= x)
	return x + y
}
