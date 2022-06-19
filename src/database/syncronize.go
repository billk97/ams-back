package database

import (
	"ams-back/src/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(models.Employee{})
			},
			Rollback: func(db *gorm.DB) error { return db.Migrator().DropTable("employees") },
		},
		{
			ID: "2",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(models.Admin{})
			},
			Rollback: func(db *gorm.DB) error { return db.Migrator().DropTable("admins") },
		},
	})
	return m.Migrate()
}
