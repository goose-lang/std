package std

import (
	"github.com/tchajed/goose/machine"
	"sync"
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

// See the [reference].
//
// [reference]: https://pkg.go.dev/bytes#Clone
func BytesClone(b []byte) []byte {
	if b == nil {
		return nil
	}
	return append([]byte{}, b...)
}

// Compute the sum of two numbers, `Assume`ing that this does not overflow.
// *Use with care*, assumptions are trusted and should be justified!
func SumAssumeNoOverflow(x uint64, y uint64) uint64 {
	machine.Assume(x+y >= x)
	return x + y
}

func Multipar(num uint64, op func(uint64)) {
	var num_left = num
	num_left_mu := new(sync.Mutex)
	num_left_cond := sync.NewCond(num_left_mu)

	for i := uint64(0); i < num; i++ {
		i := i // don't capture loop variable
		go func() {
			op(i)
			// Signal that this one is done
			num_left_mu.Lock()
			num_left -= 1
			num_left_cond.Signal()
			num_left_mu.Unlock()
		}()
	}

	num_left_mu.Lock()
	for num_left > 0 {
		num_left_cond.Wait()
	}
	num_left_mu.Unlock()
}
