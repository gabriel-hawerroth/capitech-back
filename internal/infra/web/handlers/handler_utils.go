package handlers

import "net/http"

func checkDecodeError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}
}

func setHttpError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
