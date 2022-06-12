package employee

import (
	"ams-back/src/database"
	"log"
)

func CreateEmploy(employee *Employee) *Employee {
	db := database.GetDb()
	db.AutoMigrate(&Employee{})
	err := db.Create(employee)
	if err.Error != nil {
		log.Fatal("Create employee failed")
	}
	return employee
}
