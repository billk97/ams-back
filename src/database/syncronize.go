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
		{
			ID: "4",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(models.Permission{})
			},
			Rollback: func(db *gorm.DB) error { return db.Migrator().DropTable("permissions") },
		},
		{
			ID: "5",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(models.Resource{})
			},
			Rollback: func(db *gorm.DB) error { return db.Migrator().DropTable("resources") },
		},
		{
			ID: "6",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(models.Employee{})
			},
			Rollback: func(db *gorm.DB) error { return db.Migrator().DropTable("employees") },
		},
	})
	return m.Migrate()
}
