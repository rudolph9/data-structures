package stackqueue

import "testing"

func TestStack(t *testing.T) {
	ts := Stack{}

	// Push
	ts.Push("Cecille")
	ts.Push("Kurt")
	if ts[0] != "Kurt" {
		t.Errorf("ts[0] returned: %s when it should've returned: Kurt", ts[0])
	}

	// Pop
	poppedValue := ts.Pop()
	if ts[0] != "Cecille" && poppedValue != "Kurt" {
		t.Errorf("Stack returned should be: %p, and popped value should be: %s", ts, poppedValue)
	}
}
