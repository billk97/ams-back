package models

import "gorm.io/gorm"

type VerifierToken struct {
	gorm.Model
	TokenKey string `json:"key" gorm:"size:255; UNIQUE_INDEX:invitation_index"`
	Token    string `json:"token" gorm:"size:255;"`
}
