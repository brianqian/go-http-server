package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	router  *chi.Mux
}

func New(options ...func(*Server)) *Server {
	s := &Server{host: "localhost", port: 8080, timeout: 30 * time.Second, router: chi.NewRouter()}
	for _, opt := range options {
		opt(s)
	}
	return s
}
func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}

func (s *Server) Start() {

	fmt.Println("Starting server...")

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.URLFormat)
	s.router.Use(middleware.Timeout(s.timeout))

	s.router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	s.router.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		s.respondWithJson(w, 200, struct {
			Success bool `json:"success"`
		}{true})
	})

	s.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Printf("Server running on port: %v \n", s.port)

	err := http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)

	if err != nil {
		log.Fatal(err)
	}
}
