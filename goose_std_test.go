package std

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	s := []uint64{1, 2, 3}
	s1, s2 := SliceSplit(s, 1)
	assert.Len(s1, 1)
	assert.Len(s2, 2)

	s1, s2 = SliceSplit(s, 3)
	assert.Len(s1, 3)
	assert.Len(s2, 0)
}

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
