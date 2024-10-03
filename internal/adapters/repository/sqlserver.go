package repository

import (
	"fmt"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn_sql_server = GetEnv("DB_SQLSERVER_STRING_CONECTION")
var DBSQLServer *gorm.DB

func DBConnection() {
	fmt.Println("Connecting to SQL Server database...")
	db, err := gorm.Open(sqlserver.Open(dsn_sql_server), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Dev{})

	DBSQLServer = db

	// Seed users
	var count int64
	var users []domain.User
	DBSQLServer.Find(&users).Count(&count)
	fmt.Printf("Users in database: %d\n", count)

	if count == 0 {
		SeedUsers(db)
	}
}
