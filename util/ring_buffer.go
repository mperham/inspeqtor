package util

import (
	"sort"
	"sync"
)

/*
  Metrics are stored in a RingBuffer so we can keep
  history while also keeping the total storage constant.
  You can add things to a RingBuffer but you cannot remove them.
  Older items will be overwritten and garbage collected in time.
*/
type RingBuffer struct {
	values []*float64
	oldest int
	mu     sync.Mutex
}

/*
  Return a RingBuffer with the given capacity.
*/
func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		make([]*float64, capacity), 0, sync.Mutex{},
	}
}

/*
  Add the given element to the buffer.  Method will
  panic if the caller tries to store nil.
*/
func (buf *RingBuffer) Add(elem float64) {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	idx := buf.oldest
	buf.values[idx] = &elem

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
func (buf *RingBuffer) At(idx int) *float64 {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	latest := buf.oldest + idx - 1
	if latest < 0 {
		latest = len(buf.values) + latest
	}
	return buf.values[latest]
}

func (buf *RingBuffer) Size() int {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	var count int
	for _, x := range buf.values {
		if x != nil {
			count++
		}
	}

	return count
}

func (buf *RingBuffer) Capacity() int {
	return cap(buf.values)
}

// Export the set of values in the Ring Buffer, where the latest value
// will be last in the array.  If the Buffer was not full, 0 will be exported.
func (buf *RingBuffer) Export() []float64 {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	length := len(buf.values)
	data := make([]float64, length)

	idx := 0
	for i := buf.oldest; i < len(buf.values); i++ {
		v := buf.values[i]
		if v != nil {
			data[idx] = *v
			idx++
		}
	}

	for i := 0; i < buf.oldest; i++ {
		v := buf.values[i]
		if v != nil {
			data[idx] = *v
			idx++
		}
	}

	sort.Reverse(sort.Float64Slice(data))
	shrink := make([]float64, idx)
	copy(shrink, data)

	return shrink
}
