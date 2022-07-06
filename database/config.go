package database

import (
	utils "ams-back/utils"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// TODO init a connection pull
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
	fmt.Printf(dsn)
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
