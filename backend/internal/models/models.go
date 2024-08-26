package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default:false"`
}

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Quantity    int `gorm:"not null"`
}

type Location struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Capacity int    `gorm:"not null"`
}

type Issue struct {
	gorm.Model
	ItemID   int    `gorm:"not null"`
	Item     Item   `gorm:"foreignKey:ItemID"`
	Quantity int    `gorm:"not null"`
	Status   string `gorm:"not null;default:'pending'"`
}

type InventoryReport struct {
	ID             int
	Name           string
	Quantity       int
	IssuedQuantity int
}

// If you need a custom time type for any reason, you can add this:
type CustomTime struct {
	time.Time
}
