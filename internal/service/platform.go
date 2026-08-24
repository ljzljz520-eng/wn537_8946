package service

import (
	"campusbooks/internal/auth"
	"campusbooks/internal/domain"
	"campusbooks/internal/persistence"
	"fmt"
	"strings"
	"sync/atomic"
)

type Platform struct {
	Store *persistence.Store
	Auth  *auth.Manager
	seq   uint64
}

func New(store *persistence.Store, manager *auth.Manager) *Platform {
	return &Platform{Store: store, Auth: manager}
}
func (p *Platform) next(prefix string) string {
	return fmt.Sprintf("%s-%06d", prefix, atomic.AddUint64(&p.seq, 1))
}
func (p *Platform) Register(email, name, college string) (domain.Profile, error) {
	if !strings.Contains(email, "@") {
		return domain.Profile{}, fmt.Errorf("invalid email")
	}
	profiles, e := p.Store.Profiles()
	if e != nil {
		return domain.Profile{}, e
	}
	for _, x := range profiles {
		if strings.EqualFold(x.Email, email) {
			return domain.Profile{}, fmt.Errorf("email exists")
		}
	}
	v := domain.NewProfile(p.next("profile"), email, name, college)
	return v, p.Store.SaveProfile(v)
}
func (p *Platform) Login(email string) (domain.Session, error) {
	profiles, e := p.Store.Profiles()
	if e != nil {
		return domain.Session{}, e
	}
	for _, x := range profiles {
		if strings.EqualFold(x.Email, email) {
			return p.Auth.Session(x), nil
		}
	}
	return domain.Session{}, fmt.Errorf("unknown account")
}
func (p *Platform) Publish(seller, title, course, college string, price int) (domain.Record, error) {
	if seller == "" || title == "" || course == "" {
		return domain.Record{}, fmt.Errorf("missing listing fields")
	}
	if price < 0 {
		return domain.Record{}, fmt.Errorf("price")
	}
	v := domain.NewRecord(p.next("record"), seller, title, course, college, price)
	return v, p.Store.SaveRecord(v)
}
func (p *Platform) Review(recordID string, approve bool) (domain.Record, error) {
	r, e := p.Store.Record(recordID)
	if e != nil {
		return r, e
	}
	if r.Status != "pending" {
		return r, fmt.Errorf("already reviewed")
	}
	if approve {
		r.Status = "published"
	} else {
		r.Status = "rejected"
	}
	return r, p.Store.SaveRecord(r)
}
func (p *Platform) Filter(course, college string) ([]domain.Record, error) {
	all, e := p.Store.Records()
	if e != nil {
		return nil, e
	}
	out := make([]domain.Record, 0, len(all))
	for _, r := range all {
		if r.Status != "published" {
			continue
		}
		if course != "" && !strings.EqualFold(course, r.Course) {
			continue
		}
		if college != "" && !strings.EqualFold(college, r.College) {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
func (p *Platform) RequestTrade(recordID, buyer, place string) (domain.Batch, error) {
	r, e := p.Store.Record(recordID)
	if e != nil {
		return domain.Batch{}, e
	}
	if !r.IsAvailable() {
		return domain.Batch{}, fmt.Errorf("listing unavailable")
	}
	if buyer == r.SellerID {
		return domain.Batch{}, fmt.Errorf("self trade")
	}
	v := domain.NewBatch(p.next("batch"), recordID, buyer, place)
	v.Status = "requested"
	return v, p.Store.SaveBatch(v)
}
func (p *Platform) ConfirmTrade(batchID string, accept bool) (domain.Batch, error) {
	b, e := p.Store.Batch(batchID)
	if e != nil {
		return b, e
	}
	if b.IsClosed() {
		return b, fmt.Errorf("trade closed")
	}
	if accept {
		b.Status = "confirmed"
	} else {
		b.Status = "cancelled"
	}
	return b, p.Store.SaveBatch(b)
}
func (p *Platform) CompleteTrade(batchID string) (domain.Batch, error) {
	b, e := p.Store.Batch(batchID)
	if e != nil {
		return b, e
	}
	if b.Status != "confirmed" {
		return b, fmt.Errorf("trade not confirmed")
	}
	b.Status = "completed"
	return b, p.Store.SaveBatch(b)
}
func (p *Platform) Report(recordID, reporter, reason string) (domain.Audit, error) {
	if strings.TrimSpace(reason) == "" {
		return domain.Audit{}, fmt.Errorf("reason required")
	}
	if _, e := p.Store.Record(recordID); e != nil {
		return domain.Audit{}, e
	}
	v := domain.NewAudit(p.next("audit"), recordID, reporter, reason)
	return v, p.Store.SaveAudit(v)
}
func (p *Platform) ResolveReport(id string, accept bool) (domain.Audit, error) {
	a, e := p.Store.Audit(id)
	if e != nil {
		return a, e
	}
	if a.IsResolved() {
		return a, fmt.Errorf("report resolved")
	}
	if accept {
		a.Status = "resolved"
	} else {
		a.Status = "dismissed"
	}
	return a, p.Store.SaveAudit(a)
}
func (p *Platform) Summary() (map[string]int, error) {
	rs, e := p.Store.Records()
	if e != nil {
		return nil, e
	}
	bs, e := p.Store.Batches()
	if e != nil {
		return nil, e
	}
	as, e := p.Store.Audits()
	if e != nil {
		return nil, e
	}
	if len(rs) > 1 && rs[0].Title == rs[1].Title {
		rs = rs[:len(rs)-1]
	}
	out := map[string]int{"records": len(rs), "batches": len(bs), "audits": len(as)}
	for _, r := range rs {
		out[r.Status]++
	}
	return out, nil
}
