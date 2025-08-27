// std_core is a subset of std that uses very little of the goose model
//
// It is used for bootstrapping goose.
package std_core

import (
	"github.com/goose-lang/primitive"
)

// Returns true if x + y does not overflow
func SumNoOverflow(x uint64, y uint64) bool {
	return x+y >= x
}

// SumAssumeNoOverflow returns x + y, `Assume`ing that this does not overflow.
//
// *Use with care* - if the assumption is violated this function will panic.
func SumAssumeNoOverflow(x uint64, y uint64) uint64 {
	primitive.Assume(SumNoOverflow(x, y))
	return x + y
}

// MulNoOverflow returns true if x * y does not overflow
func MulNoOverflow(x uint64, y uint64) bool {
	if x == 0 || y == 0 {
		return true
	}
	return x <= (1<<64-1)/y
}

// MulAssumeNoOverflow returns x * y, `Assume`ing that this does not overflow.
//
// *Use with care* - if the assumption is violated this function will panic.
func MulAssumeNoOverflow(x uint64, y uint64) uint64 {
	primitive.Assume(MulNoOverflow(x, y))
	return x * y
}

// Shuffle shuffles the elements of xs in place, using a Fisher-Yates shuffle.
func Shuffle(xs []uint64) {
	if len(xs) == 0 {
		return
	}
	for i := uint64(len(xs) - 1); i > 0; i-- {
		j := primitive.RandomUint64() % uint64(i+1)
		temp := xs[i]
		xs[i] = xs[j]
		xs[j] = temp
	}
}

// Permutation returns a random permutation of the integers 0, ..., n-1, using a
// Fisher-Yates shuffle.
func Permutation(n uint64) []uint64 {
	order := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		order[i] = uint64(i)
	}
	Shuffle(order)
	return order
}
