package api

import (
	"base/types"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() {
	registerUserRoutes(s.router)
}

func registerUserRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", GetUsers)
	})
}

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		userID := "THIS IS A TEST"
		// userID := chi.URLParam(r, "userId")
		ctx := context.WithValue(r.Context(), types.UserIdKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}))
}
