package domain

import "strings"

func NormalizeCollege(v string) string { return strings.ToUpper(strings.TrimSpace(v)) }
func NormalizeCourse(v string) string  { return strings.ToUpper(strings.TrimSpace(v)) }
func ValidProfile(p Profile) bool {
	return p.ID != "" && strings.Contains(p.Email, "@") && p.Name != ""
}
func ValidRecord(r Record) bool {
	return r.ID != "" && r.SellerID != "" && r.Title != "" && r.Course != "" && r.Price >= 0
}
