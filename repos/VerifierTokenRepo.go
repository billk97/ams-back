package repos

import (
	"ams-back/database"
	"ams-back/models"
)

func FindToken(key string) (*models.VerifierToken, error) {
	db := database.GetDb()
	token := models.VerifierToken{}
	result := db.Where("token_key = ?", key).First(&token)
	if result.Error != nil {
		return nil, result.Error
	}
	return &token, nil
}

func CreateToken(token *models.VerifierToken) error {
	db := database.GetDb()
	result := db.Create(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
