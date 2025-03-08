package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb(dbuser, dbpasswordv, dbhost, dbname, dbport string) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpasswordv, dbhost, dbport)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
		return nil, err
	}

	// Create database if it doesn't exist
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname).Error; err != nil {
		log.Fatalf("Error creating database: %v", err)
		return nil, err
	}

	return db, nil
}
