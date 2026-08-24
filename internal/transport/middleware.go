package transport

import (
	"net/http"
	"strings"
)

func JSONOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "json required", 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, r)
	})
}
