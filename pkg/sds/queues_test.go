package sds

import (
	"sync"
	"testing"
)

func TestBoundedQueue(t *testing.T) {
	const C = 10000
	const P = 100
	const W = C / P

	q := NewBoundedQueue[int](C)

	if cap := q.Capacity(); cap != C {
		t.Fatalf("Capacity %d, want %d\n", cap, C)
	}

	var wg sync.WaitGroup

	wg.Add(P)
	for k := 0; k < P; k++ {
		go func() {
			defer wg.Done()
			for w := 1; w <= W; w++ {
				if ok := q.Enqueue(w); !ok {
					t.Fatalf("Cannot Enqueue\n")
				}
			}
		}()
	}
	wg.Wait()

	if current := q.CurrentUse(); current != C {
		t.Fatalf("CurrentUse %d, want %d\n", current, C)
	}

	wg.Add(P)
	for k := 0; k < P; k++ {
		go func() {
			defer wg.Done()
			for w := 1; w <= W; w++ {
				if _, ok := q.Dequeue(); !ok {
					t.Fatalf("Cannot Dequeue\n")
				}
			}
		}()
	}
	wg.Wait()

	if current := q.CurrentUse(); current != 0 {
		t.Fatalf("CurrentUse %d, want %d\n", current, C)
	}

	if peak := q.PeakUse(); peak != C {
		t.Fatalf("PeakUse %d, want %d\n", peak, C)
	}
}
