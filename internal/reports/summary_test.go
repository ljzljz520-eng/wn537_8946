package reports

import (
	"campusbooks/internal/domain"
	"testing"
)

func TestBuildSnapshot(t *testing.T) {
	r := domain.NewRecord("1", "p", "T", "C", "X", 1)
	r.Status = "published"
	s := Build([]domain.Record{r}, nil, nil)
	if s.Published != 1 {
		t.Fatal(s)
	}
}
