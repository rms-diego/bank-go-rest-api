package config

import (
	"fmt"
	"os"
)

type EnvironmentVariables struct {
	ServerPort  string
	DatabaseUrl string
}

func GetEnvironmentVariables() *EnvironmentVariables {
	serverPortEnv := os.Getenv("PORT")
	databaseUrlEnv := os.Getenv("DATABASE_URL")

	// valides if environments variables was set
	if serverPortEnv == "" || databaseUrlEnv == "" {
		panic("All environment variable must be defined on '.env' file")
	}

	return &EnvironmentVariables{
		ServerPort:  fmt.Sprintf(":%v", serverPortEnv),
		DatabaseUrl: fmt.Sprintf("%s?sslmode=disable", databaseUrlEnv),
	}
}
