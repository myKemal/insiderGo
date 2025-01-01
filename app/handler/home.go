package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header and write a response
	w.Header().Set("Content-Type", "text/plain")
	_, err := fmt.Fprintln(w, "Hello! Welcome to the insiderGo server.")
	if err != nil {
		return
	}
}
