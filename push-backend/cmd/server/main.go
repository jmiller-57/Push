package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmiller-57/Push/push-backend/internal/http/router"
)

func main() {
	addr := getenv("ADDR", ":8080")
	jwtSecret := getenv("JWT_SECRET", "dev-secret-change-me")
	frontend := getenv("FRONTEND_ORIGIN", "http://localhost:5173")

	r := router.New(jwtSecret, frontend)
	log.Printf("listending on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func getenv(k, def string) string {
	if v:= os.Getenv(k); v != "" {
		return v
	}
	return def
}
