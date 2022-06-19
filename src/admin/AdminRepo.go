package admin

import (
	"ams-back/src/amserr"
	"ams-back/src/database"
	"ams-back/src/models"
	"fmt"
)

func CreateAdmin(admin *models.Admin) (*models.Admin, *amserr.ApiError) {
	db := database.GetDb()
	result := db.Create(&admin)
	if result.Error != nil {
		e := amserr.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type admin"),
		)
		return nil, e
	}
	return admin, nil
}

func FindAdminByUsername(username string) (*models.Admin, *amserr.ApiError) {
	db := database.GetDb()
	admin := models.Admin{}
	result := db.Where("username = ?", username).Find(&admin)
	if result.Error != nil {
		e := amserr.NewApiError(
			"SELECT_OPERATION_FAILED",
			result.Error,
			fmt.Sprintf("Could't find entity of type admin with username: %s", username),
		)
		return nil, e
	}
	if (admin == models.Admin{}) {
		return nil, nil
	}
	return &admin, nil
}
