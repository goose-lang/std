package std

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	Assert(true)
	assert.Panics(t, func() {
		Assert(false)
	})
}

func TestSignedSumAssumeNoOverflow(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int(3), SignedSumAssumeNoOverflow(1, 2))
	assert.Equal(int(-3), SignedSumAssumeNoOverflow(-1, -2))
	assert.Equal(int(-1), SignedSumAssumeNoOverflow(1, -2))
	assert.Equal(int(math.MaxInt), SignedSumAssumeNoOverflow(math.MaxInt, 0))
	assert.Equal(int(math.MinInt), SignedSumAssumeNoOverflow(math.MinInt, 0))
	assert.Panics(func() { SignedSumAssumeNoOverflow(math.MaxInt, 1) })
	assert.Panics(func() { SignedSumAssumeNoOverflow(math.MinInt, -1) })
	assert.Panics(func() { SignedSumAssumeNoOverflow(1, math.MaxInt) })
}

func TestBytesEqual(t *testing.T) {
	assert := assert.New(t)

	assert.True(BytesEqual([]byte{1, 2, 3}, []byte{1, 2, 3}))
	assert.False(BytesEqual([]byte{1, 2, 3}, []byte{1, 2}), "lengths differ")
	assert.False(BytesEqual([]byte{1, 3}, []byte{1, 2}), "contents differ")
}

func TestBytesEqualNilEmpty(t *testing.T) {
	assert := assert.New(t)

	// nil and empty are BytesEqual but not ==
	assert.True(BytesEqual(nil, []byte{}))
	assert.False(nil == []byte{})

	assert.False(BytesEqual(nil, []byte{1, 2}))
	assert.False(BytesEqual([]byte{}, []byte{1, 2}))
}

func TestBytesClone(t *testing.T) {
	assert := assert.New(t)

	var s0 []byte
	s1 := BytesClone(s0)
	assert.True(s1 == nil)

	s2 := []byte{1, 2}
	s3 := BytesClone(s2)
	s3[0] = 2
	assert.True(s2[0] == 1)
}

func TestSliceSplit(t *testing.T) {
	assert := assert.New(t)

	s := []byte{1, 2, 3}
	s1, s2 := SliceSplit(s, 1)
	assert.Len(s1, 1)
	assert.Len(s2, 2)

	s1, s2 = SliceSplit(s, 3)
	assert.Len(s1, 3)
	assert.Len(s2, 0)
}

func TestMultipar(t *testing.T) {
	ch := make(chan uint64)
	go Multipar(5, func(i uint64) {
		ch <- i
	})
	var results []uint64
	for i := 0; i < 5; i++ {
		results = append(results, <-ch)
	}
	assert.ElementsMatch(t, results, []uint64{0, 1, 2, 3, 4})
}

func TestSkip(t *testing.T) {
	// nothing much to test, it does nothing
	Skip()
}
