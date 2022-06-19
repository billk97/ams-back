package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Alias        string `gorm:"alias"`
	Description  string `gorm:"description"`
	Permission   Permission
	PermissionID int
}
