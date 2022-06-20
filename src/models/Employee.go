package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName       string       `json:"firstName"`
	LastName        string       `json:"lastName"`
	JobTitle        string       `json:"jobTitle"`
	BirthDate       time.Time    `json:"birthDate"`
	Email           string       `json:"email"`
	MobileNumber    string       `json:"mobileNumber"`
	DidConnectionId string       `json:"didConnectionId"`
	Permission      []Permission `json:"permission" gorm:"many2many:employee_permissions;"`
}
