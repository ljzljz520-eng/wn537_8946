package catalog

import (
	"campusbooks/internal/domain"
	"testing"
)

func TestSearchFilters(t *testing.T) {
	rs := []domain.Record{domain.NewRecord("1", "p", "Go", "CS", "Science", 5), domain.NewRecord("2", "p", "Art", "ART", "Arts", 8)}
	rs[0].Status = "published"
	if len(Search(rs, Query{Course: "CS"})) != 1 {
		t.Fatal("filter")
	}
}
