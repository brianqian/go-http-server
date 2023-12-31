package serialize

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

func Json(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON: %v", payload)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		slog.Error("Error writing")
	}
}
