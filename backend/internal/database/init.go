package database

import (
	"finance-api/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(dbuser, dbpasswordv, dbhost, dbname, dbport string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpasswordv, dbhost, dbport)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)

	}

	// Create database if it doesn't exist
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname).Error; err != nil {
		log.Fatalf("Error creating database: %v", err)

	}
	// Connect to the newly created database
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpasswordv, dbhost, dbport, dbname)
	db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)

	}
	// Set global connection
	InitDbConnection(db)

	AutoMigrateModels(db, &model.Transaction{})

	fmt.Println("Database connected and migrated successfully!")
}
