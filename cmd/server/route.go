package main

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/internal/user"
	"github.com/rms-diego/bank-go-rest-api/pkg/httpResponse"
)

func routes(appRoutes *http.ServeMux) {
	appRoutes.HandleFunc("/", healthCheck)

	user.Routes(appRoutes)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct{ Message string }

	if r.Method != http.MethodGet {
		httpResponse.NewJsonResponse(w, http.StatusBadRequest, response{Message: "route not found"})
		return
	}

	httpResponse.NewJsonResponse(w, http.StatusOK, response{Message: "Server is running"})
}
