package config

import (
	"fmt"
	"os"
)

type environmentVariables struct {
	ServerPort  string
	DatabaseUrl string
}

func GetEnvironmentVariables() *environmentVariables {
	serverPortEnv := os.Getenv("PORT")
	databaseUrlEnv := os.Getenv("DATABASE_URL")

	// valides if environments variables was set
	if serverPortEnv == "" || databaseUrlEnv == "" {
		panic("All environment variable must be defined on '.env' file")
	}

	return &environmentVariables{
		ServerPort:  fmt.Sprintf(":%v", serverPortEnv),
		DatabaseUrl: databaseUrlEnv,
	}
}
