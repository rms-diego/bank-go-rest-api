package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func routes(appRoutes *http.ServeMux) {
	appRoutes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Message string
		}

		fmt.Println(r.Body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{Message: "Server is running"})
	})

}
