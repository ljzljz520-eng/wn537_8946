package main

import (
	"campusbooks/internal/auth"
	"campusbooks/internal/persistence"
	"campusbooks/internal/service"
	"campusbooks/internal/transport"
	"log"
	"net/http"
	"time"
)

func main() {
	c := loadConfig()
	store, e := persistence.Open(c.DBPath)
	if e != nil {
		log.Fatal(e)
	}
	defer store.Close()
	p := service.New(store, auth.NewManager(c.Secret, 24*time.Hour))
	log.Println(http.ListenAndServe(c.Addr, transport.WithCORS(transport.JSONOnly(transport.New(p).Routes()))))
}
