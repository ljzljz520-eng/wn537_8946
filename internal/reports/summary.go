package reports

import (
	"campusbooks/internal/domain"
	"fmt"
	"sort"
)

type Snapshot struct{ Total, Published, Pending, Trades, OpenReports int }

func Build(records []domain.Record, batches []domain.Batch, audits []domain.Audit) Snapshot {
	s := Snapshot{Total: len(records), Trades: len(batches)}
	for _, r := range records {
		switch r.Status {
		case "published":
			s.Published++
		case "pending":
			s.Pending++
		}
	}
	for _, a := range audits {
		if a.Status == "open" {
			s.OpenReports++
		}
	}
	return s
}
func RankColleges(records []domain.Record) []string {
	counts := map[string]int{}
	for _, r := range records {
		counts[r.College]++
	}
	out := make([]string, 0, len(counts))
	for c := range counts {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return counts[out[i]] > counts[out[j]] })
	return out
}
func Describe(s Snapshot) string {
	return fmt.Sprintf("%d listings, %d published, %d pending, %d trades, %d reports", s.Total, s.Published, s.Pending, s.Trades, s.OpenReports)
}
