package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Alias string `json:"alias"`
}
