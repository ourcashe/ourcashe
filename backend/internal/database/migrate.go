package database

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Error in auto migration: %v", err)
		return err
	}
	return nil
}
