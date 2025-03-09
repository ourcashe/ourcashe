package model

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	Date        time.Time `gorm:"type:datetime;not null" json:"date"`
	Amount      float32   `gorm:"type:decimal(10,2);not null" json:"amount"`
	WhoPays     string    `gorm:"type:varchar(100);not null" json:"who_pays"`
}

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:100;unique;not null" json:"username"`
	Email     string    `gorm:"size:100;unique;not null" json:"user_email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;not null" json:"date_created"`
	IsAdmin   bool      `gorm:"default:false;not null" json:"is_admin"`
}

// Login request struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
