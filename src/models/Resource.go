package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Alias        string     `json:"alias"`
	Description  string     `json:"description"`
	Permission   Permission `json:"permission"`
	PermissionID int        `json:"permissionId" gorm:"default:Null;"`
}
