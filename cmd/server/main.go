package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rms-diego/bank-go-rest-api/pkg/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg := config.GetEnvironmentVariables()

	appRoutes := http.NewServeMux()
	routes(appRoutes)

	server := http.Server{
		Addr:    cfg.ServerPort,
		Handler: appRoutes,
	}

	fmt.Printf("\nServer is running\nLink: http://localhost%v",
		cfg.ServerPort,
	)

	server.ListenAndServe()
}
