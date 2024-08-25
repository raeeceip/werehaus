package database

import (
	"database/sql"
	"errors"
	"go-warehouse-management/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./warehouse.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS items (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            quantity INTEGER NOT NULL
        );
        CREATE TABLE IF NOT EXISTS locations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            capacity INTEGER NOT NULL
        );
        CREATE TABLE IF NOT EXISTS issues (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            item_id INTEGER,
            location_id INTEGER,
            quantity INTEGER NOT NULL,
            issue_date DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (item_id) REFERENCES items(id),
            FOREIGN KEY (location_id) REFERENCES locations(id)
        );
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL CHECK(role IN ('user', 'manager', 'admin'))
        );
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateItem(item *models.Item) error {
	_, err := db.Exec("INSERT INTO items (name, quantity) VALUES (?, ?)", item.Name, item.Quantity)
	return err
}

func CreateLocation(location *models.Location) error {
	_, err := db.Exec("INSERT INTO locations (name, capacity) VALUES (?, ?)", location.Name, location.Capacity)
	return err
}

func IssueItem(itemID, locationID, quantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if item exists and has enough quantity
	var currentQuantity int
	err = tx.QueryRow("SELECT quantity FROM items WHERE id = ?", itemID).Scan(&currentQuantity)
	if err != nil {
		return err
	}
	if currentQuantity < quantity {
		return errors.New("insufficient quantity")
	}

	// Update item quantity
	_, err = tx.Exec("UPDATE items SET quantity = quantity - ? WHERE id = ?", quantity, itemID)
	if err != nil {
		return err
	}

	// Create issue record
	_, err = tx.Exec("INSERT INTO issues (item_id, location_id, quantity) VALUES (?, ?, ?)", itemID, locationID, quantity)
	if err != nil {
		return err
	}

	return tx.Commit()
}
