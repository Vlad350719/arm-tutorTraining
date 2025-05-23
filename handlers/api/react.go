package api

import (
	"encoding/json"
	"net/http"
)

func ReactHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "tendons and joints are the most important!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
