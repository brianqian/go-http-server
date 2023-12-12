package api

import (
	"base/handlers"
	s "base/pkg/serializers"
	"net/http"
)

var userService handlers.UserService

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := userService.FindUserById(r.Context(), "fake_id")
	s.Json(w, 200, user)
}
