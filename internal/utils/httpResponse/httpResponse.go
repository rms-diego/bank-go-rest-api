package httpResponse

import (
	"encoding/json"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func NewJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func NewErrorResponse(w http.ResponseWriter, status int, err error) {

	r := errResponse{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}
