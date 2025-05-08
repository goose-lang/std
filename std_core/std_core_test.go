package std_core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNoOverflow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, SumNoOverflow(2, 3))
	assert.Equal(true, SumNoOverflow(1<<64-2, 1))
	assert.Equal(true, SumNoOverflow(1, 1<<64-2))

	assert.Equal(false, SumNoOverflow(1<<64-1, 1))
	assert.Equal(false, SumNoOverflow(1<<63, 1<<63))
	assert.Equal(false, SumNoOverflow(1<<64-7, 26))
}

func TestSumAssumeNoOverflow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uint64(5), SumAssumeNoOverflow(2, 3))
	assert.Equal(uint64(5), SumAssumeNoOverflow(5, 0))
	assert.Equal(uint64(1<<64-1), SumAssumeNoOverflow(1<<64-2, 1))
	assert.Panics(func() {
		SumAssumeNoOverflow(1<<63, 1<<63)
	})
	assert.Panics(func() {
		SumAssumeNoOverflow(1<<64-1, 2)
	})
}

func TestMulAssumeNoOverflow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uint64(6), MulAssumeNoOverflow(2, 3))
	assert.Equal(uint64(0), MulAssumeNoOverflow(0, 3))
	assert.Equal(uint64(1<<64-1), MulAssumeNoOverflow(1<<32-1, 1<<32+1))
	assert.Equal(uint64(1<<64-1), MulAssumeNoOverflow(1<<32+1, 1<<32-1))
	assert.Panics(func() {
		MulAssumeNoOverflow(1<<63, 2)
	})
	assert.Panics(func() {
		MulAssumeNoOverflow(2, 1<<63)
	})
}

func TestPermutation(t *testing.T) {
	assert := assert.New(t)

	order := Permutation(1)
	assert.ElementsMatch(order, []uint64{0})

	order = Permutation(2)
	assert.ElementsMatch(order, []uint64{0, 1})

	order = Permutation(5)
	assert.ElementsMatch(order, []uint64{0, 1, 2, 3, 4})
}

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	xs := []uint64{}
	Shuffle(xs)

	xs = []uint64{1, 1, 1}
	Shuffle(xs)
	assert.Equal(xs, []uint64{1, 1, 1})
}
