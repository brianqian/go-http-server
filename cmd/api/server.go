package api

import (
	"base/internal/db"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
	router *chi.Mux
	db     *db.Database
}

func Main(db *db.Database) {
	s := InitServer()
	s.db = db
	s.RegisterRoutes()

	fmt.Printf("Server running on %s \n", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
