package repository

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	variable := os.Getenv(key)

	if variable == "" && key == "PORT" {
		return "5000"
	}

	if variable == "" {
		fmt.Println("The environment variable " + key + " is not set. Using default value.")
	}

	return variable
}
