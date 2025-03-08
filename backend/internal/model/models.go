package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
	Username  string    `gorm:"size:100;unique;not null" json:"name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;not null" json:"date"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
