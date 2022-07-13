package repos

import (
	"ams-back/database"
	"ams-back/models"
	"ams-back/utils"
	"github.com/google/uuid"
)

func SaveEmploy(employee *models.Employee) (*models.Employee, error) {

	// TODO convert it to usecase
	db := database.GetDb()
	employee.Invitation = uuid.New().String()
	result := db.Create(employee)
	if result.Error != nil {
		return nil, result.Error
	}
	utils.SendEmail(employee.Invitation, employee.Email)
	return employee, nil
}

// todo add repository call get all rooms a user can access

func UpdateEmployee(employee *models.Employee) (*models.Employee, error) {
	db := database.GetDb()
	result := db.Save(employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil
}

func FindEmployeeById(id int) (*models.Employee, error) {
	db := database.GetDb()
	var employee models.Employee
	result := db.Preload("Permission").First(&employee, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func FindEmployeeByInvitation(invitation string) (*models.Employee, error) {
	db := database.GetDb()
	employee := models.Employee{}
	result := db.Where("invitation = ?", invitation).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func FindEmployees(page int) (*[]models.Employee, error) {
	offset := 0
	if page > 0 {
		offset = page - 1
	}
	db := database.GetDb()
	var employees []models.Employee
	result := db.
		Offset(offset).
		Limit(20).
		Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employees, nil
}
