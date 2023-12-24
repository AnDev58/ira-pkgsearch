package util

import (
	"mime"
	"net/http"
)

func EnforceJSON(w http.ResponseWriter, r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")

	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return true
	}
	return false
}
