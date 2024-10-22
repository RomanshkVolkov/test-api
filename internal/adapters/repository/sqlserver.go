package repository

import (
	"fmt"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var dsn_sql_server = GetEnv("DB_SQLSERVER_STRING_CONECTION")

var DSNList, _ = GetDSNList()
var DBSQLServer []DSNSource

func DBConnection() {
	fmt.Println("DSNList: ", DSNList)
	for _, dsn := range DSNList {
		db, err := gorm.Open(sqlserver.Open(dsn.DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			panic("failed to connect database")
		}

		DBSQLServer = append(DBSQLServer, DSNSource{
			Name: dsn.Name,
			DB:   db,
		})

		db.AutoMigrate(&domain.User{})
		db.AutoMigrate(&domain.Dev{})
		db.AutoMigrate(&domain.UserProfiles{})

		// Seed users
		var count int64
		var profiles []domain.UserProfiles
		db.Find(&profiles).Count(&count)

		if count == 0 {
			fmt.Println("Seeding user profiles")
			SeedProfiles(db)
		}

		var users []domain.User
		db.Find(&users).Count(&count)
		if count == 0 {
			fmt.Println("Seeding users")
			SeedUsers(db)
		}

	}

}

func GetDBConnection(subdomain string) DSNSource {
	fmt.Println("Subdomain: ", subdomain)
	authorizedHost := MAPPED_AUTHORIZED_DOMAINS[subdomain]
	for _, db := range DBSQLServer {
		if db.Name == authorizedHost {
			fmt.Println("DB: ", db.Name)
			return db
		}
	}

	return DSNSource{}
}
