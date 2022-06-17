package employee

import (
	"ams-back/src/amserr"
	"ams-back/src/database"
	"fmt"
	"log"
)

func CreateEmploy(employee *Employee) *Employee {
	db := database.GetDb()
	db.AutoMigrate(&Employee{})
	result := db.Create(&employee)
	if result.Error != nil {
		log.Fatal("Create employee failed")
	}
	return employee
}

func FindEmployeeById(id int) (*Employee, *amserr.ApiError) {
	db := database.GetDb()
	var employee Employee
	result := db.First(&employee, id)
	if result.Error != nil {
		e := amserr.NewApiError(
			"NOT_FOUND",
			result.Error,
			fmt.Sprintf("Could't find employee withId: %d", id),
		)
		return nil, e
	}
	return &employee, nil
}
