package database

import (
	"database/sql"
	"fmt"
	"go-warehouse-management/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./warehouse.db")
	if err != nil {
		return err
	}

	// Create tables if they don't exist
	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	// Create items table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			quantity INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Create locations table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS locations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			capacity INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Create issues table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS issues (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			item_id INTEGER,
			quantity INTEGER NOT NULL,
			from_location_id INTEGER,
			to_location_id INTEGER,
			status TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (item_id) REFERENCES items (id),
			FOREIGN KEY (from_location_id) REFERENCES locations (id),
			FOREIGN KEY (to_location_id) REFERENCES locations (id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func GetItems(page, limit int, search string) ([]models.Item, error) {
	query := "SELECT id, name, description, quantity FROM items"
	args := []interface{}{}

	if search != "" {
		query += " WHERE name LIKE ?"
		args = append(args, "%"+search+"%")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func CreateItem(item *models.Item) error {
	result, err := db.Exec("INSERT INTO items (name, description, quantity) VALUES (?, ?, ?)",
		item.Name, item.Description, item.Quantity)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.ID = fmt.Sprintf("%d", id)
	return nil
}

func UpdateItem(item *models.Item) error {
	_, err := db.Exec("UPDATE items SET name = ?, description = ?, quantity = ? WHERE id = ?",
		item.Name, item.Description, item.Quantity, item.ID)
	return err
}

func DeleteItem(id string) error {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}

func GetLocations() ([]models.Location, error) {
	rows, err := db.Query("SELECT id, name, capacity FROM locations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.ID, &location.Name, &location.Capacity)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func CreateLocation(location *models.Location) error {
	result, err := db.Exec("INSERT INTO locations (name, capacity) VALUES (?, ?)",
		location.Name, location.Capacity)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	location.ID = fmt.Sprintf("%d", id)
	return nil
}

func UpdateLocation(location *models.Location) error {
	_, err := db.Exec("UPDATE locations SET name = ?, capacity = ? WHERE id = ?",
		location.Name, location.Capacity, location.ID)
	return err
}

func DeleteLocation(id string) error {
	_, err := db.Exec("DELETE FROM locations WHERE id = ?", id)
	return err
}

func CreateIssue(issue *models.Issue) error {
	result, err := db.Exec("INSERT INTO issues (item_id, quantity, from_location_id, to_location_id, status) VALUES (?, ?, ?, ?, ?)",
		issue.ItemID, issue.Quantity, issue.FromLocationID, issue.ToLocationID, "pending")
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	issue.ID = fmt.Sprintf("%d", id)
	return nil
}

func GetPendingIssues() ([]models.Issue, error) {
	rows, err := db.Query(`
		SELECT i.id, i.item_id, it.name, i.quantity, i.from_location_id, fl.name, i.to_location_id, tl.name, i.created_at
		FROM issues i
		JOIN items it ON i.item_id = it.id
		JOIN locations fl ON i.from_location_id = fl.id
		JOIN locations tl ON i.to_location_id = tl.id
		WHERE i.status = 'pending'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []models.Issue
	for rows.Next() {
		var issue models.Issue
		err := rows.Scan(&issue.ID, &issue.ItemID, &issue.ItemName, &issue.Quantity,
			&issue.FromLocationID, &issue.FromLocationName, &issue.ToLocationID, &issue.ToLocationName,
			&issue.CreatedAt)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, nil
}

func ApproveIssue(id string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var issue models.Issue
	err = tx.QueryRow("SELECT item_id, quantity, from_location_id, to_location_id FROM issues WHERE id = ?", id).
		Scan(&issue.ItemID, &issue.Quantity, &issue.FromLocationID, &issue.ToLocationID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE items SET quantity = quantity - ? WHERE id = ?", issue.Quantity, issue.ItemID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE issues SET status = 'approved' WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func DenyIssue(id string) error {
	_, err := db.Exec("UPDATE issues SET status = 'denied' WHERE id = ?", id)
	return err
}

func GetInventoryReport() ([]models.InventoryReport, error) {
	rows, err := db.Query(`
		SELECT i.id, i.name, i.quantity, COALESCE(SUM(iss.quantity), 0) as issued_quantity
		FROM items i
		LEFT JOIN issues iss ON i.id = iss.item_id AND iss.status = 'approved'
		GROUP BY i.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []models.InventoryReport
	for rows.Next() {
		var item models.InventoryReport
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.IssuedQuantity)
		if err != nil {
			return nil, err
		}
		report = append(report, item)
	}

	return report, nil
}

func GetIssueReport() ([]models.IssueReport, error) {
	rows, err := db.Query(`
		SELECT i.id, it.name, i.quantity, fl.name, tl.name, i.status, i.created_at
		FROM issues i
		JOIN items it ON i.item_id = it.id
		JOIN locations fl ON i.from_location_id = fl.id
		JOIN locations tl ON i.to_location_id = tl.id
		ORDER BY i.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []models.IssueReport
	for rows.Next() {
		var issue models.IssueReport
		err := rows.Scan(&issue.ID, &issue.ItemName, &issue.Quantity, &issue.FromLocation,
			&issue.ToLocation, &issue.Status, &issue.CreatedAt)
		if err != nil {
			return nil, err
		}
		report = append(report, issue)
	}

	return report, nil
}

func GetItemMovementReport(itemID string) ([]models.ItemMovement, error) {
	rows, err := db.Query(`
		SELECT i.id, i.quantity, fl.name, tl.name, i.status, i.created_at
		FROM issues i
		JOIN locations fl ON i.from_location_id = fl.id
		JOIN locations tl ON i.to_location_id = tl.id
		WHERE i.item_id = ?
		ORDER BY i.created_at DESC
	`, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []models.ItemMovement
	for rows.Next() {
		var movement models.ItemMovement
		err := rows.Scan(&movement.IssueID, &movement.Quantity, &movement.FromLocation,
			&movement.ToLocation, &movement.Status, &movement.CreatedAt)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil
}
