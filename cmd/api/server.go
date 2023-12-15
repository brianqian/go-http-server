package api

import (
	"base/internal/db"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	server *http.Server
	router *chi.Mux
	db     *pgxpool.Pool
}

func Main() {
	s := InitServer()
	s.db = db.New(db.DbConfig{MinConnections: "1", MaxConnections: "4"})
	s.RegisterRoutes()

	fmt.Printf("Server running on %s \n", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
