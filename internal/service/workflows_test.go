package service

import (
	"campusbooks/internal/auth"
	"campusbooks/internal/persistence"
	"path/filepath"
	"testing"
	"time"
)

func testPlatform(t *testing.T) *Platform {
	t.Helper()
	s, e := persistence.Open(filepath.Join(t.TempDir(), "data.db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return New(s, auth.NewManager("test-secret", time.Hour))
}
func TestWorkflowAccept(t *testing.T) {
	p := testPlatform(t)
	id, e := p.WorkflowAccept("a@school.edu", "A", "Science", "Go Basics", "CS101")
	if e != nil || id == "" {
		t.Fatalf("workflow: %v", e)
	}
	r, e := p.Store.Record(id)
	if e != nil || r.Status != "published" {
		t.Fatalf("record: %+v %v", r, e)
	}
}
func TestWorkflowPublish(t *testing.T) {
	p := testPlatform(t)
	id, e := p.WorkflowAccept("a@school.edu", "A", "Science", "Go Basics", "CS101")
	if e != nil {
		t.Fatal(e)
	}
	buyer, _ := p.Register("b@school.edu", "B", "Science")
	bid, e := p.WorkflowPublish(id, buyer.ID, "Library")
	if e != nil || bid == "" {
		t.Fatalf("publish: %v", e)
	}
}
func TestWorkflowReopen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "reopen.db")
	s, e := persistence.Open(path)
	if e != nil {
		t.Fatal(e)
	}
	p := New(s, auth.NewManager("x", time.Hour))
	u, _ := p.Register("a@x.edu", "A", "Arts")
	p.Publish(u.ID, "Book", "HIS1", "Arts", 5)
	s.Close()
	s, e = persistence.Open(path)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	rs, e := s.Records()
	if e != nil || len(rs) != 1 {
		t.Fatalf("reopen: %d %v", len(rs), e)
	}
}
func TestPersistenceSurvivesReopen(t *testing.T) { TestWorkflowReopen(t) }
func TestWorkflow11(t *testing.T) {
	p := testPlatform(t)
	u, _ := p.Register("a@x.edu", "A", "Arts")
	if _, e := p.Publish(u.ID, "Same", "HIS1", "Arts", 5); e != nil {
		t.Fatal(e)
	}
	if _, e := p.Publish(u.ID, "Same", "HIS1", "Arts", 5); e != nil {
		t.Fatal(e)
	}
	sum, e := p.Summary()
	if e != nil {
		t.Fatal(e)
	}
	if sum["records"] != 2 {
		t.Fatalf("duplicate registrations lost: %+v", sum)
	}
}
