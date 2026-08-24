package persistence

import (
	"campusbooks/internal/domain"
	"path/filepath"
	"testing"
)

func TestCountAll(t *testing.T) {
	s, _ := Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveProfile(domain.NewProfile("p", "p@x", "P", "C"))
	m, e := s.CountAll()
	if e != nil || m["profiles"] != 1 {
		t.Fatalf("%v %v", m, e)
	}
}
