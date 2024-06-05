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

	config := config.GetEnvironmentVariables()

	serverRunningMessage := fmt.Sprintf("Server is running\nLink: http://localhost:%v", 3000)
	fmt.Println(serverRunningMessage)
	http.ListenAndServe(config.ServerPort, nil)
}
