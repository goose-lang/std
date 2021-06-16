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
