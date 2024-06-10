package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/config"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg := config.GetEnvironmentVariables()

	err = database.Connect(cfg)

	if err != nil {
		panic(err)
	}

	appRoutes := http.NewServeMux()
	routes(appRoutes)

	server := http.Server{
		Addr:    cfg.ServerPort,
		Handler: appRoutes,
	}

	fmt.Printf("Server is running\nLink: http://localhost%v\n\n",
		cfg.ServerPort,
	)

	server.ListenAndServe()
}
