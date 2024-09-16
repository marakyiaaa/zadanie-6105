package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress    string
	PostgresConn     string
	PostgresJDBCURL  string
	PostgresUsername string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
}

func MustConfig() *Config {
	envVars := map[string]string{
		"SERVER_ADDRESS":    os.Getenv("SERVER_ADDRESS"),
		"POSTGRES_CONN":     os.Getenv("POSTGRES_CONN"),
		"POSTGRES_JDBC_URL": os.Getenv("POSTGRES_JDBC_URL"),
		"POSTGRES_USERNAME": os.Getenv("POSTGRES_USERNAME"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_HOST":     os.Getenv("POSTGRES_HOST"),
		"POSTGRES_PORT":     os.Getenv("POSTGRES_PORT"),
		"POSTGRES_DATABASE": os.Getenv("POSTGRES_DATABASE"),
	}

	for envVar, value := range envVars {
		if value == "" {
			log.Fatalf("%s environment variable is not set", envVar)
		}
	}

	return &Config{
		ServerAddress:    os.Getenv("SERVER_ADDRESS"),
		PostgresConn:     os.Getenv("POSTGRES_CONN"),
		PostgresJDBCURL:  os.Getenv("POSTGRES_JDBC_URL"),
		PostgresUsername: os.Getenv("POSTGRES_USERNAME"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
	}
}
