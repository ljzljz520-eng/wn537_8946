package auth

import (
	"campusbooks/internal/domain"
	"testing"
	"time"
)

func TestTokenRoundTrip(t *testing.T) {
	m := NewManager("secret", time.Hour)
	tok, e := m.Issue(domain.NewProfile("p", "p@x", "P", "C"))
	if e != nil {
		t.Fatal(e)
	}
	id, e := m.Verify(tok)
	if e != nil || id != "p" {
		t.Fatalf("%s %v", id, e)
	}
}
