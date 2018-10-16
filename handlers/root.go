package handlers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	msg := "{\"greeting\":\"hello\"}"
	writeBytes(w, []byte(msg))
}
