package auth

import (
	"campusbooks/internal/domain"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type Manager struct {
	secret []byte
	ttl    time.Duration
}

func NewManager(secret string, ttl time.Duration) *Manager { return &Manager{[]byte(secret), ttl} }
func (m *Manager) Issue(p domain.Profile) (string, error) {
	if len(m.secret) == 0 {
		return "", fmt.Errorf("secret required")
	}
	exp := time.Now().Add(m.ttl).Unix()
	raw := p.ID + "." + fmt.Sprint(exp)
	mac := hmac.New(sha256.New, m.secret)
	mac.Write([]byte(raw))
	return base64.RawURLEncoding.EncodeToString([]byte(raw)) + "." + base64.RawURLEncoding.EncodeToString(mac.Sum(nil)), nil
}
func (m *Manager) Verify(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("malformed token")
	}
	raw, e := base64.RawURLEncoding.DecodeString(parts[0])
	if e != nil {
		return "", e
	}
	sig, e := base64.RawURLEncoding.DecodeString(parts[1])
	if e != nil {
		return "", e
	}
	mac := hmac.New(sha256.New, m.secret)
	mac.Write(raw)
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return "", fmt.Errorf("signature mismatch")
	}
	fields := strings.Split(string(raw), ".")
	if len(fields) != 2 {
		return "", fmt.Errorf("claims")
	}
	var exp int64
	if _, e = fmt.Sscan(fields[1], &exp); e != nil {
		return "", e
	}
	if time.Now().Unix() > exp {
		return "", fmt.Errorf("expired")
	}
	return fields[0], nil
}
func (m *Manager) Session(p domain.Profile) domain.Session {
	t, _ := m.Issue(p)
	return domain.Session{Token: t, ProfileID: p.ID, ExpiresAt: time.Now().Add(m.ttl)}
}
