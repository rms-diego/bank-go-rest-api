package config

import (
	"fmt"
	"os"
)

type EnvironmentVariables struct {
	ServerPort  string
	DatabaseUrl string
	JwtSecret   string
}

func GetEnvironmentVariables() *EnvironmentVariables {
	serverPortEnv := os.Getenv("PORT")
	databaseUrlEnv := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	// valides if environments variables was set
	if serverPortEnv == "" || databaseUrlEnv == "" || jwtSecret == "" {
		panic("All environment variable must be defined on '.env' file")
	}

	return &EnvironmentVariables{
		ServerPort:  fmt.Sprintf(":%v", serverPortEnv),
		DatabaseUrl: fmt.Sprintf("%s?sslmode=disable", databaseUrlEnv),
		JwtSecret:   jwtSecret,
	}
}
