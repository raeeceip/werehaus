package models

import "time"

type Item struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

type Location struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Issue struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint
	ItemID         uint
	FromLocationID uint
	ToLocationID   uint
	Quantity       int
	Status         string
	IssueDate      time.Time
	ApprovalDate   *time.Time
	ApprovedBy     *uint
	User           User     `gorm:"foreignKey:UserID"`
	Item           Item     `gorm:"foreignKey:ItemID"`
	FromLocation   Location `gorm:"foreignKey:FromLocationID"`
	ToLocation     Location `gorm:"foreignKey:ToLocationID"`
	ApprovedByUser *User    `gorm:"foreignKey:ApprovedBy"`
}

type InventoryReport struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"`
	IssuedQuantity int    `json:"issued_quantity"`
}

type IssueReport struct {
	ID           uint      `json:"id"`
	ItemName     string    `json:"item_name"`
	Quantity     int       `json:"quantity"`
	FromLocation string    `json:"from_location"`
	ToLocation   string    `json:"to_location"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type ItemMovement struct {
	IssueID      uint      `json:"issue_id"`
	Quantity     int       `json:"quantity"`
	FromLocation string    `json:"from_location"`
	ToLocation   string    `json:"to_location"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
