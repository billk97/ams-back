package repos

import (
	"ams-back/database"
	"ams-back/models"
	"log"
)

func CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	db := database.GetDb()
	result := db.Create(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return admin, nil
}

func FindAdminByUsername(username string) (*models.Admin, error) {
	db := database.GetDb()
	admin := models.Admin{}
	result := db.Where("username = ?", username).Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	if (admin == models.Admin{}) {
		log.Printf("Entity admin with username: %s not found", username)
		return nil, nil
	}
	return &admin, nil
}
