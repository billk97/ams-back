package employee

import (
	"ams-back/src/amserr"
	"ams-back/src/database"
	"fmt"
	"log"
)

func SaveEmploy(employee *Employee) *Employee {
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

func FindEmployees(page int) (*[]Employee, *amserr.ApiError) {
	offset := 0
	if page > 0 {
		offset = page - 1
	}
	db := database.GetDb()
	var employees []Employee
	result := db.
		Offset(offset).
		Limit(20).
		Find(&employees)
	if result.Error != nil {
		e := amserr.NewApiError(
			"NOT_FOUND",
			result.Error,
			fmt.Sprintf("Could't find employees: "),
		)
		return nil, e
	}
	return &employees, nil
}
