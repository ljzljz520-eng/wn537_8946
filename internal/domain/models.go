package domain

import "time"

type Profile struct {
	ID, Email, Name, College string
	CreatedAt                time.Time
}
type Record struct {
	ID, SellerID, Title, Course, College, Status string
	Price                                        int
	CreatedAt                                    time.Time
}
type Batch struct {
	ID, RecordID, BuyerID, MeetingPlace, Status string
	CreatedAt                                   time.Time
}
type Audit struct {
	ID, RecordID, ReporterID, Reason, Status string
	CreatedAt                                time.Time
}
type Session struct {
	Token, ProfileID string
	ExpiresAt        time.Time
}

func NewProfile(id, email, name, college string) Profile {
	return Profile{ID: id, Email: email, Name: name, College: college, CreatedAt: time.Now().UTC()}
}
func NewRecord(id, seller, title, course, college string, price int) Record {
	return Record{ID: id, SellerID: seller, Title: title, Course: course, College: college, Price: price, Status: "pending", CreatedAt: time.Now().UTC()}
}
func NewBatch(id, record, buyer, place string) Batch {
	return Batch{ID: id, RecordID: record, BuyerID: buyer, MeetingPlace: place, Status: "requested", CreatedAt: time.Now().UTC()}
}
func NewAudit(id, record, reporter, reason string) Audit {
	return Audit{ID: id, RecordID: record, ReporterID: reporter, Reason: reason, Status: "open", CreatedAt: time.Now().UTC()}
}
func (r Record) IsAvailable() bool { return r.Status == "published" }
func (b Batch) IsClosed() bool     { return b.Status == "completed" || b.Status == "cancelled" }
func (a Audit) IsResolved() bool   { return a.Status == "resolved" || a.Status == "dismissed" }
