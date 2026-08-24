package service

import (
	"campusbooks/internal/domain"
	"fmt"
	"sort"
	"strings"
	"time"
)

type ListingView struct {
	Record   domain.Record
	Seller   domain.Profile
	CanTrade bool
	Age      time.Duration
}
type TradeView struct {
	Batch  domain.Batch
	Record domain.Record
	Buyer  domain.Profile
	Seller domain.Profile
}

func (p *Platform) ListingView(id string) (ListingView, error) {
	r, e := p.Store.Record(id)
	if e != nil {
		return ListingView{}, e
	}
	u, e := p.Store.Profile(r.SellerID)
	if e != nil {
		return ListingView{}, e
	}
	return ListingView{Record: r, Seller: u, CanTrade: r.IsAvailable(), Age: time.Since(r.CreatedAt)}, nil
}
func (p *Platform) ListingViews(q string) ([]ListingView, error) {
	rs, e := p.Store.Records()
	if e != nil {
		return nil, e
	}
	out := []ListingView{}
	for _, r := range rs {
		if q != "" && !strings.Contains(strings.ToLower(r.Title), strings.ToLower(q)) {
			continue
		}
		v, e := p.ListingView(r.ID)
		if e == nil {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Record.CreatedAt.After(out[j].Record.CreatedAt) })
	return out, nil
}
func (p *Platform) TradeView(id string) (TradeView, error) {
	b, e := p.Store.Batch(id)
	if e != nil {
		return TradeView{}, e
	}
	r, e := p.Store.Record(b.RecordID)
	if e != nil {
		return TradeView{}, e
	}
	u, e := p.Store.Profile(b.BuyerID)
	if e != nil {
		return TradeView{}, e
	}
	seller, e := p.Store.Profile(r.SellerID)
	if e != nil {
		return TradeView{}, e
	}
	return TradeView{Batch: b, Record: r, Buyer: u, Seller: seller}, nil
}
func (p *Platform) TradesForStudent(id string) ([]TradeView, error) {
	bs, e := p.Store.Batches()
	if e != nil {
		return nil, e
	}
	out := []TradeView{}
	for _, b := range bs {
		r, e := p.Store.Record(b.RecordID)
		if e != nil {
			continue
		}
		if b.BuyerID != id && r.SellerID != id {
			continue
		}
		v, e := p.TradeView(b.ID)
		if e == nil {
			out = append(out, v)
		}
	}
	return out, nil
}
func (p *Platform) CancelTrade(id string) (domain.Batch, error) {
	b, e := p.Store.Batch(id)
	if e != nil {
		return b, e
	}
	if b.Status == "completed" {
		return b, fmt.Errorf("completed trade")
	}
	if b.Status == "cancelled" {
		return b, fmt.Errorf("already cancelled")
	}
	b.Status = "cancelled"
	return b, p.Store.SaveBatch(b)
}
func (p *Platform) MoveMeeting(id, place string) (domain.Batch, error) {
	b, e := p.Store.Batch(id)
	if e != nil {
		return b, e
	}
	if b.IsClosed() {
		return b, fmt.Errorf("closed trade")
	}
	if strings.TrimSpace(place) == "" {
		return b, fmt.Errorf("meeting place required")
	}
	b.MeetingPlace = strings.TrimSpace(place)
	return b, p.Store.SaveBatch(b)
}
func (p *Platform) ChangePrice(id string, price int) (domain.Record, error) {
	r, e := p.Store.Record(id)
	if e != nil {
		return r, e
	}
	if price < 0 {
		return r, fmt.Errorf("negative price")
	}
	if r.Status == "archived" || r.Status == "rejected" {
		return r, fmt.Errorf("inactive listing")
	}
	r.Price = price
	return r, p.Store.SaveRecord(r)
}
func (p *Platform) Relist(id string) (domain.Record, error) {
	r, e := p.Store.Record(id)
	if e != nil {
		return r, e
	}
	if r.Status != "archived" && r.Status != "rejected" {
		return r, fmt.Errorf("cannot relist")
	}
	r.Status = "pending"
	return r, p.Store.SaveRecord(r)
}
func (p *Platform) Archive(id string) (domain.Record, error) {
	r, e := p.Store.Record(id)
	if e != nil {
		return r, e
	}
	if r.Status == "archived" {
		return r, fmt.Errorf("already archived")
	}
	r.Status = "archived"
	return r, p.Store.SaveRecord(r)
}
func (p *Platform) StudentListings(id string) ([]domain.Record, error) {
	rs, e := p.Store.Records()
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range rs {
		if r.SellerID == id {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out, nil
}
func (p *Platform) StudentReports(id string) ([]domain.Audit, error) {
	as, e := p.Store.Audits()
	if e != nil {
		return nil, e
	}
	out := []domain.Audit{}
	for _, a := range as {
		if a.ReporterID == id {
			out = append(out, a)
		}
	}
	return out, nil
}
func (p *Platform) ReportCount(id string) (int, error) {
	as, e := p.Store.Audits()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, a := range as {
		if a.RecordID == id {
			n++
		}
	}
	return n, nil
}
func (p *Platform) IsTrusted(id string) bool {
	as, e := p.StudentReports(id)
	if e != nil {
		return false
	}
	open := 0
	for _, a := range as {
		if a.Status == "open" {
			open++
		}
	}
	return open < 3
}
func (p *Platform) Dashboard(id string) (map[string]int, error) {
	ls, e := p.StudentListings(id)
	if e != nil {
		return nil, e
	}
	ts, e := p.TradesForStudent(id)
	if e != nil {
		return nil, e
	}
	rs, e := p.StudentReports(id)
	if e != nil {
		return nil, e
	}
	out := map[string]int{"listings": len(ls), "trades": len(ts), "reports": len(rs)}
	for _, r := range ls {
		if r.Status == "published" {
			out["active"]++
		}
		if r.Status == "pending" {
			out["awaiting_review"]++
		}
	}
	return out, nil
}
func (p *Platform) ValidateOwnership(id, student string) error {
	r, e := p.Store.Record(id)
	if e != nil {
		return e
	}
	if r.SellerID != student {
		return fmt.Errorf("student does not own listing")
	}
	return nil
}
func (p *Platform) DeleteListing(id, student string) error {
	if e := p.ValidateOwnership(id, student); e != nil {
		return e
	}
	r, e := p.Store.Record(id)
	if e != nil {
		return e
	}
	if r.Status == "published" {
		return fmt.Errorf("unpublish before delete")
	}
	_, e = p.Archive(id)
	return e
}
func (p *Platform) EligibleBuyer(recordID, buyer string) bool {
	r, e := p.Store.Record(recordID)
	if e != nil {
		return false
	}
	return r.IsAvailable() && r.SellerID != buyer
}
func (p *Platform) SearchForStudent(student, course, college string) ([]domain.Record, error) {
	rs, e := p.Filter(course, college)
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range rs {
		if r.SellerID != student {
			out = append(out, r)
		}
	}
	return out, nil
}
func (p *Platform) MarkViewed(student, record string) error {
	if student == "" || record == "" {
		return fmt.Errorf("view requires ids")
	}
	if _, e := p.Store.Record(record); e != nil {
		return e
	}
	return nil
}
func (p *Platform) Health() map[string]string {
	return map[string]string{"status": "ok", "service": "campusbooks", "time": time.Now().UTC().Format(time.RFC3339)}
}
