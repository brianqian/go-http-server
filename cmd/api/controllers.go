package api

import (
	s "base/cmd/services"
	"base/internal/serialize"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	userService  = s.UserService{}
	chessService = s.ChessService{}
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

func GetChessComPgnByMonth(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	date := chi.URLParam(r, "date")

	res, err := chessService.GetChessComPgn(r.Context(), username, date)

	if err != nil {
		slog.Error("Errorsdfas")
	}

	if len(res) == 0 {
		fmt.Println(res)
		serialize.Json(w, 200, res)
		return
	}
	serialize.Json(w, 200, res)

}
