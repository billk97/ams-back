package database

import (
	"ams-back/src/utils"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDbConnection(config utils.DB) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=%t&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		true,
		"Local")
	dbc, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Database connection error: ", err)
		return
	}
	db = dbc
	log.Println("Database connection established")
}

func GetDb() *gorm.DB {
	return db
}
