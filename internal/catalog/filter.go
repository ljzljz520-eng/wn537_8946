package catalog

import (
	"campusbooks/internal/domain"
	"sort"
	"strings"
)

type Query struct {
	Course, College, Text string
	MaxPrice              int
}

func Match(r domain.Record, q Query) bool {
	if r.Status != "published" {
		return false
	}
	if q.Course != "" && !strings.EqualFold(r.Course, q.Course) {
		return false
	}
	if q.College != "" && !strings.EqualFold(r.College, q.College) {
		return false
	}
	if q.Text != "" && !strings.Contains(strings.ToLower(r.Title), strings.ToLower(q.Text)) {
		return false
	}
	return q.MaxPrice <= 0 || r.Price <= q.MaxPrice
}
func Search(records []domain.Record, q Query) []domain.Record {
	out := make([]domain.Record, 0)
	for _, r := range records {
		if Match(r, q) {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out
}
func Colleges(records []domain.Record) []string {
	set := map[string]bool{}
	for _, r := range records {
		if r.College != "" {
			set[r.College] = true
		}
	}
	out := []string{}
	for c := range set {
		out = append(out, c)
	}
	sort.Strings(out)
	return out
}
func Courses(records []domain.Record) []string {
	set := map[string]bool{}
	for _, r := range records {
		if r.Course != "" {
			set[r.Course] = true
		}
	}
	out := []string{}
	for c := range set {
		out = append(out, c)
	}
	sort.Strings(out)
	return out
}
