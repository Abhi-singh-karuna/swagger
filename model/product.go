package model

import "github.com/jinzhu/gorm"

// Product struct
type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}

// User struct
type User struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
	//	Amount      int    `gorm:"not null" json:"amount"`
}
