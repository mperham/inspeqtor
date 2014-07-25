package util

import (
	"sync"
)

/*
  Inspeqtor metrics are stored in a RingBuffer so we can keep
  history while also keeping the total storage constant.
  You can add things to this RingBuffer but you cannot remove them.
  Older items will be overwritten and garbage collected in time.
*/
type RingBuffer struct {
	values []interface{}
	oldest int
	mu     sync.Mutex
}

/*
  Return a RingBuffer with the given capacity.
*/
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		make([]interface{}, size), 0, sync.Mutex{},
	}
}

/*
  Add the given element to the buffer.  Method will
  panic if the caller tries to store nil.
*/
func (buf *RingBuffer) Add(elem interface{}) {
	if elem == nil {
		panic("Attempting to store a nil value")
	}

	buf.mu.Lock()
	defer buf.mu.Unlock()

	idx := buf.oldest
	buf.values[idx] = elem

	buf.oldest = idx + 1
	if buf.oldest >= len(buf.values) {
		buf.oldest = 0
	}
}

/*
  Access the ring buffer based on previous elements added.
  0 is the latest item, -1 is the previous one added, etc.
  This means that At(1) will give you the oldest item.
  Returns nil if the slot in the buffer has not been filled
  yet.
*/
func (buf *RingBuffer) At(idx int) interface{} {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	latest := buf.oldest + idx - 1
	if latest < 0 {
		latest = len(buf.values) + latest
	}
	return buf.values[latest]
}
