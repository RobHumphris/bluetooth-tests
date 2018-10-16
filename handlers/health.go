package handlers

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	stats := Stats.Data()
	b, _ := json.Marshal(stats)
	writeBytes(w, b)
}
