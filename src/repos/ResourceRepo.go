package repos

import (
	"ams-back/src/database"
	"ams-back/src/models"
	"ams-back/src/utils"
	"fmt"
)

func SaveResource(resource *models.Resource) *utils.ApiError {
	db := database.GetDb()
	result := db.Create(resource)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type Resource"),
		)
		return e
	}
	return nil
}

func UpdateResource(resource *models.Resource) *utils.ApiError {
	db := database.GetDb()
	result := db.Save(resource)
	if result.Error != nil {
		e := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			result.Error,
			fmt.Sprintf("Could't persist entity of type Resource"),
		)
		return e
	}
	return nil
}

func FindOneResourceById(id int) (*models.Resource, *utils.ApiError) {
	db := database.GetDb()
	resource := &models.Resource{}
	result := db.Where("id = ?", id).First(&resource)
	if result.Error != nil {
		e := utils.NewApiError(
			"QUERY_EXECUTION_FAILED",
			result.Error,
			fmt.Sprintf("Could not execute query"),
		)
		return nil, e
	}
	if (resource == &models.Resource{}) {
		return nil, nil
	}
	return resource, nil
}

func FindResource(page int) (*[]models.Resource, *utils.ApiError) {
	offset := 0
	db := database.GetDb()
	resources := []models.Resource{}
	if page > 0 {
		offset = page - 1
	}
	result := db.Offset(offset).Limit(20).Find(&resources)
	if result.Error != nil {
		e := utils.NewApiError(
			"QUERY_EXECUTION_FAILED",
			result.Error,
			fmt.Sprintf("Could not execute query"),
		)
		return nil, e
	}
	return &resources, nil
}
