package database

import (
	"ams-back/utils"
	"fmt"
	"log"
	"time"

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
	sqlDb, err := dbc.DB()
	if err != nil {
		log.Fatalf(err.Error())
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	db = dbc
	log.Println("Database connection established")
}

func GetDb() *gorm.DB {
	return db
}
