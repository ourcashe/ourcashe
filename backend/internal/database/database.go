package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DbConnection *gorm.DB // Global DB instance

func InitDbConnection(db *gorm.DB) {
	DbConnection = db
}

func InsertRecord(record interface{}) error {
	if DbConnection == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	result := DbConnection.Create(record)
	fmt.Println(result)
	if result.Error != nil {
		fmt.Println("something went wrong in db query")
		// handle error here
		return result.Error
	}
	fmt.Println("record inserted successfully ", result.RowsAffected)
	return nil
}

func GetAllRecords(record interface{}) error {
	if DbConnection == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	result := DbConnection.Find(record)
	if result.Error != nil {
		fmt.Println("something went wrong in db query")
		// handle error here
		return result.Error
	}
	fmt.Println("record get successfully ", result.RowsAffected)
	return nil
}

// DeleteRecordByID deletes a record from the database based on model and ID
func DeleteRecordByID(model interface{}, id interface{}) error {
	if DbConnection == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// Delete using ID
	result := DbConnection.Delete(model, id)
	if result.Error != nil {
		fmt.Println("Error in DB query:", result.Error)
		return result.Error
	}

	// Check if a record was actually deleted
	if result.RowsAffected == 0 {
		fmt.Println("No record found to delete")
		return fmt.Errorf("record not found")
	}

	fmt.Println("Record deleted successfully")
	return nil
}
