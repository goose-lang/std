package std

import (
	"sync"

	"github.com/goose-lang/goose/machine"
)

// BytesEqual returns if the two byte slices are equal.
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

// SliceSplit splits xs at n into two slices.
//
// The capacity of the first slice overlaps with the second, so afterward it is
// no longer safe to append to the first slice.
//
// TODO: once goose supports it, make this function generic in the slice element
// type
func SliceSplit(xs []byte, n uint64) ([]byte, []byte) {
	// TODO: we could get ownership of xs's capacity if we could write xs[:n:n]
	// (this would reset xs to have no extra capacity), but Goose doesn't
	// support that.
	return xs[:n], xs[n:]
}

// Returns true if x + y does not overflow
func SumNoOverflow(x uint64, y uint64) bool {
	return x+y >= x
}

// SumAssumeNoOverflow returns x + y, `Assume`ing that this does not overflow.
//
// *Use with care* - if the assumption is violated this function will panic.
func SumAssumeNoOverflow(x uint64, y uint64) uint64 {
	machine.Assume(SumNoOverflow(x, y))
	return x + y
}

// Multipar runs op(0) ... op(num-1) in parallel and waits for them all to finish.
//
// Implementation note: does not use a done channel (which is the standard
// pattern in Go) because this is not supported by Goose. Instead uses mutexes
// and condition variables since these are modeled in Goose
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
