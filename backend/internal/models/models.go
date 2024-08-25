package models

import "time"

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

type Location struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type Issue struct {
	ID               string    `json:"id"`
	ItemID           string    `json:"item_id"`
	ItemName         string    `json:"item_name"`
	Quantity         int       `json:"quantity"`
	FromLocationID   string    `json:"from_location_id"`
	FromLocationName string    `json:"from_location_name"`
	ToLocationID     string    `json:"to_location_id"`
	ToLocationName   string    `json:"to_location_name"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type InventoryReport struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"`
	IssuedQuantity int    `json:"issued_quantity"`
}

type IssueReport struct {
	ID           string    `json:"id"`
	ItemName     string    `json:"item_name"`
	Quantity     int       `json:"quantity"`
	FromLocation string    `json:"from_location"`
	ToLocation   string    `json:"to_location"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type ItemMovement struct {
	IssueID      string    `json:"issue_id"`
	Quantity     int       `json:"quantity"`
	FromLocation string    `json:"from_location"`
	ToLocation   string    `json:"to_location"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
