package controllers

import (
	"net/http"

	s "github.com/brianqian/go-http-server/cmd/webapp-api/services"
)

var userService = s.UserService{}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := userService.FindUserById(r.Context(), "fake_id")
	RespondWithJson(w, 200, user)
}
