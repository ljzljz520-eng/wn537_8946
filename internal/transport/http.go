package transport

import (
	"campusbooks/internal/catalog"
	"campusbooks/internal/service"
	"encoding/json"
	"net/http"
)

type Server struct{ Platform *service.Platform }

func New(p *service.Platform) *Server { return &Server{Platform: p} }
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", s.health)
	mux.HandleFunc("/api/register", s.register)
	mux.HandleFunc("/api/listings", s.listings)
	mux.HandleFunc("/api/search", s.search)
	mux.HandleFunc("/api/summary", s.summary)
	return mux
}
func write(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	write(w, map[string]string{"status": "ok"})
}
func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var in struct{ Email, Name, College string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "bad request", 400)
		return
	}
	p, e := s.Platform.Register(in.Email, in.Name, in.College)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	write(w, p)
}
func (s *Server) listings(w http.ResponseWriter, r *http.Request) {
	records, e := s.Platform.Filter("", "")
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	write(w, records)
}
func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	records, _ := s.Platform.Store.Records()
	q := catalog.Query{Course: r.URL.Query().Get("course"), College: r.URL.Query().Get("college"), Text: r.URL.Query().Get("q")}
	write(w, catalog.Search(records, q))
}
func (s *Server) summary(w http.ResponseWriter, r *http.Request) {
	v, e := s.Platform.Summary()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	write(w, v)
}
