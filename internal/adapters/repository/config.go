package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"gorm.io/gorm"
)

type DSNSource struct {
	Name string
	DB   *gorm.DB
}

var MAPPED_AUTHORIZED_DOMAINS = map[string]string{
	"localhost:8080": "DB_DSN_DOMAIN_1",
	"0.0.0.0:8080":   "DB_DSN_DOMAIN_2",
	"127.0.0.1:8080": "DB_DSN_DOMAIN_3",
}

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

func GetDSNList() ([]domain.DSNSourceConfig, error) {
	availableEnvs := os.Environ()

	if len(availableEnvs) == 0 {
		fmt.Println("No environment variables available.")
		return nil, fmt.Errorf("no environment variables available")
	}

	var dsnList []domain.DSNSourceConfig

	for _, env := range availableEnvs {
		if strings.Contains(env, "DB_DSN_") {
			split := strings.Split(env, "=")
			name, dsn := split[0], strings.Join(split[1:], "=")

			dsnList = append(dsnList, domain.DSNSourceConfig{
				Name: name,
				DSN:  dsn,
			})
		}
	}

	fmt.Println("===================== DEBUG DSNList =====================")
	fmt.Println("DSNList: ", dsnList)

	return dsnList, nil

}
