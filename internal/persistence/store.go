package persistence

import (
	"campusbooks/internal/domain"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"sync"
	"time"
)

var buckets = map[string][]byte{"profiles": []byte("profiles"), "records": []byte("records"), "batches": []byte("batches"), "audits": []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put[T any](s *Store, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](s *Store, b []byte, key string) (T, error) {
	var out T
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func list[T any](s *Store, b []byte) ([]T, error) {
	out := []T{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(b).ForEach(func(_, v []byte) error {
			var x T
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveProfile(v domain.Profile) error { return put(s, buckets["profiles"], v.ID, v) }
func (s *Store) Profile(id string) (domain.Profile, error) {
	return get[domain.Profile](s, buckets["profiles"], id)
}
func (s *Store) Profiles() ([]domain.Profile, error) {
	return list[domain.Profile](s, buckets["profiles"])
}
func (s *Store) SaveRecord(v domain.Record) error { return put(s, buckets["records"], v.ID, v) }
func (s *Store) Record(id string) (domain.Record, error) {
	return get[domain.Record](s, buckets["records"], id)
}
func (s *Store) Records() ([]domain.Record, error) { return list[domain.Record](s, buckets["records"]) }
func (s *Store) SaveBatch(v domain.Batch) error    { return put(s, buckets["batches"], v.ID, v) }
func (s *Store) Batch(id string) (domain.Batch, error) {
	return get[domain.Batch](s, buckets["batches"], id)
}
func (s *Store) Batches() ([]domain.Batch, error) { return list[domain.Batch](s, buckets["batches"]) }
func (s *Store) SaveAudit(v domain.Audit) error   { return put(s, buckets["audits"], v.ID, v) }
func (s *Store) Audit(id string) (domain.Audit, error) {
	return get[domain.Audit](s, buckets["audits"], id)
}
func (s *Store) Audits() ([]domain.Audit, error) { return list[domain.Audit](s, buckets["audits"]) }
