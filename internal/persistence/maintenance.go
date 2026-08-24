package persistence

import (
	"campusbooks/internal/domain"
	"go.etcd.io/bbolt"
)

func (s *Store) SeedProfile(p domain.Profile) error {
	if !domain.ValidProfile(p) {
		return nil
	}
	return s.SaveProfile(p)
}
func (s *Store) CountAll() (map[string]int, error) {
	ps, e := s.Profiles()
	if e != nil {
		return nil, e
	}
	rs, e := s.Records()
	if e != nil {
		return nil, e
	}
	bs, e := s.Batches()
	if e != nil {
		return nil, e
	}
	as, e := s.Audits()
	if e != nil {
		return nil, e
	}
	return map[string]int{"profiles": len(ps), "records": len(rs), "batches": len(bs), "audits": len(as)}, nil
}
func (s *Store) ClearRecords() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.DeleteBucket(buckets["records"]) })
}
