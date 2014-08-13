package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	t.Parallel()
	rb := NewRingBuffer(5)

	assert.Panics(t, func() {
		rb.Add(nil)
	})
	assert.Nil(t, rb.At(0))

	rb.Add(1)

	val := rb.At(0)
	assert.Equal(t, 1, val)
	assert.Nil(t, rb.At(1))

	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	assert.Nil(t, rb.At(1))

	rb.Add(5)

	// Now that the buffer is full, we can access the oldest element
	val = rb.At(1)
	assert.Equal(t, 1, val)

	rb.Add(6)

	val = rb.At(0)
	assert.Equal(t, 6, val)
	val = rb.At(-1)
	assert.Equal(t, 5, val)
	val = rb.At(-2)
	assert.Equal(t, 4, val)
	val = rb.At(-3)
	assert.Equal(t, 3, val)
	val = rb.At(-4)
	assert.Equal(t, 2, val)
	val = rb.At(-5)
	assert.Equal(t, 6, val)
}
