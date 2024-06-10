package main

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/internal/auth"
	"github.com/rms-diego/bank-go-rest-api/internal/user"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/httpResponse"
)

func routes(appRoutes *http.ServeMux) {
	appRoutes.HandleFunc("/", healthCheck)

	user.Routes(appRoutes)
	auth.Routes(appRoutes)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct{ Message string }

	if r.Method != http.MethodGet {
		httpResponse.NewJsonResponse(w, http.StatusNotImplemented, response{Message: "not implemented"})
		return
	}

	httpResponse.NewJsonResponse(w, http.StatusOK, response{Message: "Server is running"})
}
