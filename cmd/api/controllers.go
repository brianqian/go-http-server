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
	userId := chi.URLParam(r, "userId")
	user := userService.FindUserById(r.Context(), userId)
	serialize.Json(w, 200, user)
}

func GetChessProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user := chessService.GetProfileByUsername(r.Context(), username)
	serialize.Json(w, 200, user)
}
