package repository

import (
	"fmt"

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

		RunSeeds(db)
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
