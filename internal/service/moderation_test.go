package service

import "testing"

func TestModerationRejectsUnknown(t *testing.T) {
	p := testPlatform(t)
	u, _ := p.Register("m@x.edu", "M", "Law")
	r, _ := p.Publish(u.ID, "Book", "L1", "Law", 2)
	if _, e := p.Moderate(r.ID, "wat"); e == nil {
		t.Fatal("expected action error")
	}
}
