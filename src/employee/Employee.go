package employee

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName       string
	LastName        string
	JobTitle        string
	DirthDate       time.Time
	Email           string
	MobileNumber    string
	DidConnectionId string
}
