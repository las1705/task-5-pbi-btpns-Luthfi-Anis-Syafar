package database

import (
	"UserPhoto-API/config/db_config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetDatabase() {
	var errConnection error

	if db_config.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)

		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	}

	if errConnection != nil {
		panic("cant connext to database")
	}
	log.Printf("connected to database")
}
