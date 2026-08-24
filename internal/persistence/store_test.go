package persistence

import (
	"campusbooks/internal/domain"
	"path/filepath"
	"testing"
)

func TestStoreEntities(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.SaveProfile(domain.NewProfile("p", "p@x", "P", "C")); e != nil {
		t.Fatal(e)
	}
	if e = s.SaveRecord(domain.NewRecord("r", "p", "T", "C1", "C", 1)); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Profile("p"); e != nil {
		t.Fatal(e)
	}
}
