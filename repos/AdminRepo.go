package repos

import (
	database "ams-back/database"
	models "ams-back/models"
	utils "ams-back/utils"
	"fmt"
	"log"
)

func CreateAdmin(admin *models.Admin) (*models.Admin, *utils.ApiError) {
	db := database.GetDb()
	result := db.Create(&admin)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type admin"),
		)
		return nil, e
	}
	return admin, nil
}

func FindAdminByUsername(username string) (*models.Admin, *utils.ApiError) {
	db := database.GetDb()
	admin := models.Admin{}
	result := db.Where("username = ?", username).Find(&admin)
	if result.Error != nil {
		e := utils.NewApiError(
			"SELECT_OPERATION_FAILED",
			result.Error,
			fmt.Sprintf("Could't find entity of type admin with username: %s", username),
		)
		return nil, e
	}
	if (admin == models.Admin{}) {
		log.Printf("Entity admin with username: %s not found", username)
		return nil, nil
	}
	return &admin, nil
}
