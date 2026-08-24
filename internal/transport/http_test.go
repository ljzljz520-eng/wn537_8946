package transport

import (
	"campusbooks/internal/auth"
	"campusbooks/internal/persistence"
	"campusbooks/internal/service"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
)

func TestHealth(t *testing.T) {
	s, _ := persistence.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	h := New(service.New(s, auth.NewManager("x", time.Hour))).Routes()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest("GET", "/api/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
