package api

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() {
	s.router.Route("/api", func(r chi.Router) {
		registerUserRoutes(r)
		registerChessRoutes(r)
	})
}

func registerUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Route("/{userId}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Get("/", GetUsers)
		})
	})
}
func registerChessRoutes(r chi.Router) {
	r.Route("/chess", func(r chi.Router) {
		r.Get("/profile/{username}", GetChessProfile)
	})
}
