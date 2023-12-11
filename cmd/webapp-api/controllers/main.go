package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON: %v", payload)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
