package main

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/pkg/httpResponse"
)

func routes(appRoutes *http.ServeMux) {
	appRoutes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Message string
		}

		httpResponse.NewJsonResponse(w, http.StatusOK, response{Message: "Server is running"})
	})

}
