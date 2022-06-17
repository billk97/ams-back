package employee

import (
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

func FindEmployeeById(id int) (*Employee, *ApiError) {
	db := database.GetDb()
	var employee Employee
	result := db.First(&employee, id)
	if result.Error != nil {
		// log.Fatalf("Could No find employee withId: %d", id)
		// log.Printf("Could No find employee withId: %d", id)
		e := ApiError{}
		e.Err = fmt.Sprintf("%s", result.Error)
		e.Details = fmt.Sprintf("Could No find employee withId: %d", id)
		e.init()
		return nil, &e
	}
	return &employee, nil
}
