package main

import "os"

type Config struct{ DBPath, Secret, Addr string }

func loadConfig() Config {
	c := Config{DBPath: "campusbooks.db", Secret: "campusbooks-development-secret", Addr: ":8080"}
	if v := os.Getenv("CAMPUSBOOKS_DB"); v != "" {
		c.DBPath = v
	}
	if v := os.Getenv("CAMPUSBOOKS_SECRET"); v != "" {
		c.Secret = v
	}
	if v := os.Getenv("CAMPUSBOOKS_ADDR"); v != "" {
		c.Addr = v
	}
	return c
}
