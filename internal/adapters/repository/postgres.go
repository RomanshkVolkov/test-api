package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn_postgres = GetEnv("DB_POSTGRES_STRING_CONECTION")
var DBPostgres *gorm.DB

func DBConnectionPOSTGRES() {
	db, err := gorm.Open(postgres.Open(dsn_postgres), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DBPostgres = db
}
