package api

import (
	"base/internal/serialize"
	"base/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	userService  services.UserService
	chessService services.ChessService
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := userService.FindUserById(r.Context(), "fake_id")
	serialize.Json(w, 200, user)
}

func GetChessProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user := chessService.GetProfileByUsername(r.Context(), username)
	serialize.Json(w, 200, user)
}
