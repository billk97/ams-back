package usecases

import (
	"ams-back/models"
	"ams-back/repos"
	"ams-back/utils"
	"errors"
	"github.com/google/uuid"
)

func CreateEmployeeIfNotExists(employee *models.Employee) (uint, error) {
	err := checkIfEmployeeWithIdExists(int(employee.ID))
	if err != nil {
		return 0, err
	}
	employee.Invitation = uuid.New().String()
	id, err := repos.SaveEmploy(employee)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("something when wrong during persistence phase")
	}
	emailErr := utils.SendEmail(employee.Invitation, employee.Email)
	if emailErr != nil {
		return 0, emailErr
	}
	return id, nil
}

func checkIfEmployeeWithIdExists(id int) error {
	if id == 0 {
		return nil
	}
	_, err := repos.FindEmployeeById(id)
	if err != nil {
		return err
	}
	return nil
}
