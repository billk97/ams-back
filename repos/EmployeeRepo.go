package repos

import (
	database "ams-back/database"
	models "ams-back/models"
	utils "ams-back/utils"
	"fmt"
	"github.com/google/uuid"
)

func SaveEmploy(employee *models.Employee) (*models.Employee, *utils.ApiError) {

	db := database.GetDb()
	employee.Invitation = uuid.New().String()
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

func FindEmployeeByInvitation(invitation string) (*models.Employee, *utils.ApiError) {
	db := database.GetDb()
	employee := models.Employee{}
	result := db.Where("invitation = ?", invitation).First(&employee)
	if result.Error != nil {
		fmt.Println(result.Error)
		e := utils.NewApiError(
			"NOT_FOUND",
			result.Error,
			fmt.Sprintf("Could't find employee with invitaion: %s", invitation),
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
