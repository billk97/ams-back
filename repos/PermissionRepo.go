package repos

import (
	"ams-back/database"
	"ams-back/models"
)

func FindPermissions(page int) (*[]models.Permission, error) {
	offset := 0
	if page > 0 {
		offset = page - 1
	}
	permissions := []models.Permission{}
	db := database.GetDb()
	result := db.Offset(offset).Limit(20).Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &permissions, nil
}

func CreatePermission(permission *models.Permission) error {
	db := database.GetDb()
	result := db.Create(permission)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
