package httpResponse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func NewErrorResponse(w http.ResponseWriter, status int, err error) {
	type errorMessage struct {
		Error string `json:"error"`
	}

	fmt.Printf("\nerror: %v\n", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage{Error: err.Error()})
}
