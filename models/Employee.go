package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName       string       `json:"firstName" gorm:"size:255;"`
	LastName        string       `json:"lastName" gorm:"size:255;"`
	JobTitle        string       `json:"jobTitle" gorm:"size:255;"`
	Email           string       `json:"email" gorm:"size:255;UNIQUE_INDEX:invitation_index"`
	Invitation      string       `json:"invitation" gorm:"size:255;UNIQUE_INDEX:invitation_index;"`
	DidConnectionId string       `json:"didConnectionId" gorm:"size:255;"`
	Did             string       `json:"did" gorm:"size:255;"`
	Status          string       `json:"status" gorm:"size:255;"`
	Permission      []Permission `json:"permissions" gorm:"many2many:employee_permissions;"`

	// TODO add a state if a credential has been issued?
}
