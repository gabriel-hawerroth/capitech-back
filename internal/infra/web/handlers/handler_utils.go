package handlers

import "net/http"

const (
	errorDecodingRequestBody = "Error decoding request body"
	errorCastingParams       = "Error casting params"
	errorEncodingResponse    = "Error encoding response"
)

func setHttpError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
