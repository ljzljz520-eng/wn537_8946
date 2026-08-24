package domain

import "testing"

func TestEntityStates(t *testing.T) {
	r := NewRecord("r", "p", "T", "C", "X", 1)
	if r.IsAvailable() {
		t.Fatal("pending available")
	}
	r.Status = "published"
	if !r.IsAvailable() {
		t.Fatal("published unavailable")
	}
}
