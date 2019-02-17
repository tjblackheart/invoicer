package handlers

import (
	"net/http"
)

// Status returns a pong for a ping request
func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\"pong\""))
}
