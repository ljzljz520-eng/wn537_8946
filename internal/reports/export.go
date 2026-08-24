package reports

import (
	"campusbooks/internal/domain"
	"encoding/json"
)

func JSON(records []domain.Record) ([]byte, error) {
	return json.Marshal(map[string]any{"listings": records, "count": len(records)})
}
func StatusCounts(records []domain.Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.Status]++
	}
	return out
}
