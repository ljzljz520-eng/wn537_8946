package service

import (
	"campusbooks/internal/domain"
	"fmt"
	"strings"
)

func (p *Platform) Moderate(recordID, action string) (domain.Record, error) {
	r, e := p.Store.Record(recordID)
	if e != nil {
		return r, e
	}
	switch strings.ToLower(action) {
	case "approve":
		r.Status = "published"
	case "reject":
		r.Status = "rejected"
	case "archive":
		r.Status = "archived"
	default:
		return r, fmt.Errorf("unknown moderation action")
	}
	return r, p.Store.SaveRecord(r)
}
func (p *Platform) Pending() ([]domain.Record, error) {
	all, e := p.Store.Records()
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range all {
		if r.Status == "pending" {
			out = append(out, r)
		}
	}
	return out, nil
}
