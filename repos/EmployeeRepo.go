package repos

import (
	"ams-back/database"
	"ams-back/models"
)

func SaveEmploy(employee *models.Employee) (uint, error) {
	db := database.GetDb()
	result := db.Create(employee)
	if result.Error != nil {
		return 0, result.Error
	}
	return employee.ID, nil
}

func GetEmployeeResources(id int) (*[]models.Resource, error) {
	var resources *[]models.Resource
	db := database.GetDb()
	result := db.
		Joins("JOIN employee_permissions ep on ep.permission_id = resources.permission_id").
		Where("ep.employee_id=?", id).
		Find(&resources)
	if result.Error != nil {
		return nil, result.Error
	}
	return resources, nil

}

func UpdateEmployee(employee *models.Employee) (*models.Employee, error) {
	db := database.GetDb()
	db.Model(employee).Association("Permission").Replace(employee.Permission)
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
