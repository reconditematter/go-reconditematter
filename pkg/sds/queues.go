package sds

import (
	"sync"
)

// The type Queue specifies a first-in, first-out, goroutine-safe queue.
type Queue[T any] interface {
	// Enqueue appends the element `e` at the tail of the queue.
	// If the queue is bounded, and the number of existing elements
	// equals the capacity, then Enqueue returns false without
	// modifiying the queue; otherwise Enqueue returns true.
	Enqueue(e T) bool

	// Dequeue removes the element at the head of the queue.
	// If the queue is empty, then Dequeue returns (zero,false) without
	// modifying the queue; otherwise Dequeue returns (element,true).
	Dequeue() (T, bool)

	// Capacity returns the maximum number of elements that are
	// allowed in the bounded queue. If the queue is unbounded,
	// then Capacity returns -1.
	Capacity() int

	// CurrentUse returns the number of elements currently in the queue.
	CurrentUse() int

	// PeakUse returns the maximum number of elements that have been in the queue.
	PeakUse() int
}

var _ Queue[struct{}] = &boundedQueue[struct{}]{}

type boundedQueue[T any] struct {
	lock *sync.Mutex
	elem []T
	head int
	peak int
	size int
}

// NewBoundedQueue returns a bounded queue that allows
// the maximum of `capacity` elements.
func NewBoundedQueue[T any](capacity int) Queue[T] {
	return &boundedQueue[T]{elem: make([]T, capacity)}
}

func (q *boundedQueue[T]) Enqueue(e T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == len(q.elem) {
		return false
	}

	i := q.head + q.size
	if i >= len(q.elem) {
		i -= len(q.elem)
	}
	q.elem[i] = e
	q.size++
	if q.size > q.peak {
		q.peak = q.size
	}
	return true
}

func (q *boundedQueue[T]) Dequeue() (T, bool) {
	var zero T
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return zero, false
	}

	r := q.elem[q.head]
	q.elem[q.head] = zero
	q.size--
	q.head++
	if q.head == len(q.elem) {
		q.head = 0
	}
	return r, true
}

func (q *boundedQueue[T]) Capacity() (capacity int) {
	// no need to lock here for a bounded queue
	return len(q.elem)
}

func (q *boundedQueue[T]) CurrentUse() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size
}

func (q *boundedQueue[T]) PeakUse() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.peak
}
