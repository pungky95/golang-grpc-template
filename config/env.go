package config

import (
	"fmt"
	"os"
)

var requiredEnvVars = []string{
	"APP_ENV",
	"PORT",
	"DB_HOST",
	"DB_PORT",
	"DB_DATABASE",
	"DB_USERNAME",
	"DB_PASSWORD",
	"API_KEY",
	"DEFAULT_USER_PASSWORD",
}

func CheckEnv() error {
	for _, envVar := range requiredEnvVars {
		if GetEnv(envVar, "") == "" {
			return fmt.Errorf("%s is required", envVar)
		}
	}
	return nil
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
