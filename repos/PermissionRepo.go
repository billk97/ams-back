package repos

import (
	database "ams-back/database"
	models "ams-back/models"
	utils "ams-back/utils"
	"fmt"
)

func FindPermissions(page int) (*[]models.Permission, *utils.ApiError) {
	offset := 0
	if page > 0 {
		offset = page - 1
	}
	permissions := []models.Permission{}
	db := database.GetDb()
	result := db.Offset(offset).Limit(20).Find(&permissions)
	if result.Error != nil {
		e := utils.NewApiError(
			"QUERY_EXECUTION_FAILED",
			result.Error,
			fmt.Sprintf("Could not execute query"),
		)
		return nil, e
	}
	return &permissions, nil
}

func CreatePermission(permission *models.Permission) *utils.ApiError {
	db := database.GetDb()
	result := db.Create(permission)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type permision"),
		)
		return e
	}
	return nil
}
