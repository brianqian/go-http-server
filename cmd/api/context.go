package api

import (
	"base/types"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userId")
		ctx := context.WithValue(r.Context(), types.UserIdKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}))
}
