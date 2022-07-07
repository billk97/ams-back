package repos

import (
	"ams-back/database"
	"ams-back/models"
)

func SaveResource(resource *models.Resource) error {
	db := database.GetDb()
	result := db.Create(resource)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateResource(resource *models.Resource) error {
	db := database.GetDb()
	result := db.Save(&resource)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindOneResourceById(id int) (*models.Resource, error) {
	db := database.GetDb()
	resource := &models.Resource{}
	result := db.Where("id = ?", id).First(&resource)
	if result.Error != nil {
		return nil, result.Error
	}
	if (resource == &models.Resource{}) {
		return nil, nil
	}
	return resource, nil
}

func FindResource(page int) (*[]models.Resource, error) {
	offset := 0
	db := database.GetDb()
	resources := []models.Resource{}
	if page > 0 {
		offset = page - 1
	}
	result := db.Offset(offset).Limit(20).Find(&resources)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resources, nil
}
