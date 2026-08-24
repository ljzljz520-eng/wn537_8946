package catalog

import "campusbooks/internal/domain"

func Page(records []domain.Record, page, size int) []domain.Record {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	start := (page - 1) * size
	if start >= len(records) {
		return []domain.Record{}
	}
	end := start + size
	if end > len(records) {
		end = len(records)
	}
	return records[start:end]
}
func PriceRange(records []domain.Record) (int, int) {
	if len(records) == 0 {
		return 0, 0
	}
	min, max := records[0].Price, records[0].Price
	for _, r := range records[1:] {
		if r.Price < min {
			min = r.Price
		}
		if r.Price > max {
			max = r.Price
		}
	}
	return min, max
}
