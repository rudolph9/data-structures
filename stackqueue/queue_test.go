package stackqueue

import "testing"

func TestQueue(t *testing.T) {
	tq := Queue{}

	// Enqueue
	tq.Enqueue("Cecille")
	tq.Enqueue("Kurt")
	if tq[0] != "Kurt" {
		t.Errorf("tq[0] returned: %s when it should've returned: Kurt", tq[0])
	}

	// Dequeue
	dequeuedValue := tq.Dequeue()
	if tq[0] != "Kurt" && dequeuedValue != "Cecille" {
		t.Errorf("Queue returned should be: %p, and popped value should be: %s", tq, dequeuedValue)
	}
}
