package service

import "campusbooks/internal/persistence"

func (p *Platform) WorkflowAccept(email, name, college, title, course string) (string, error) {
	u, e := p.Register(email, name, college)
	if e != nil {
		return "", e
	}
	r, e := p.Publish(u.ID, title, course, college, 10)
	if e != nil {
		return "", e
	}
	if _, e = p.Review(r.ID, true); e != nil {
		return "", e
	}
	return r.ID, nil
}
func (p *Platform) WorkflowPublish(recordID, buyer, place string) (string, error) {
	b, e := p.RequestTrade(recordID, buyer, place)
	if e != nil {
		return "", e
	}
	b, e = p.ConfirmTrade(b.ID, true)
	if e != nil {
		return "", e
	}
	b, e = p.CompleteTrade(b.ID)
	if e != nil {
		return "", e
	}
	return b.ID, nil
}
func (p *Platform) WorkflowReopen(path string) (int, error) {
	s, e := persistence.Open(path)
	if e != nil {
		return 0, e
	}
	defer s.Close()
	rs, e := s.Records()
	if e != nil {
		return 0, e
	}
	return len(rs), nil
}
