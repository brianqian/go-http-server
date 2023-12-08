package controllers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working")
	w.Write([]byte("User"))
}
