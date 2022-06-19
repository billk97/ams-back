package repos

import (
	"ams-back/src/database"
	"ams-back/src/models"
	"ams-back/src/utils"
	"fmt"
)

func SaveEmploy(employee *models.Employee) (*models.Employee, *utils.ApiError) {
	db := database.GetDb()
	fmt.Printf("%+v", employee)
	result := db.Create(employee)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type employee"),
		)
		return nil, e
	}
	return employee, nil
}

func UpdateEmployee(employee *models.Employee) (*models.Employee, error) {
	db := database.GetDb()
	result := db.Save(employee)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type employee"),
		)
		return nil, e
	}
	return employee, nil
}

func FindEmployeeById(id int) (*models.Employee, *utils.ApiError) {
	db := database.GetDb()
	var employee models.Employee
	result := db.First(&employee, id)
	if result.Error != nil {
		e := utils.NewApiError(
			"NOT_FOUND",
			result.Error,
			fmt.Sprintf("Could't find employee withId: %d", id),
		)
		return nil, e
	}
	return &employee, nil
}

func FindEmployees(page int) (*[]models.Employee, *utils.ApiError) {
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
		e := utils.NewApiError(
			"NOT_FOUND",
			result.Error,
			fmt.Sprintf("Could't find employees: "),
		)
		return nil, e
	}
	return &employees, nil
}
