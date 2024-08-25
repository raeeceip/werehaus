package database

import (
	"go-warehouse-management/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("warehouse.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Location{}, &models.Issue{})
	if err != nil {
		return err
	}

	return PopulateWithSampleData()
}

func ResetDatabase() error {
	// Drop all tables
	err := DB.Migrator().DropTable(&models.User{}, &models.Item{}, &models.Location{}, &models.Issue{})
	if err != nil {
		return err
	}

	// Recreate tables
	err = DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Location{}, &models.Issue{})
	if err != nil {
		return err
	}

	return PopulateWithSampleData()
}

func PopulateWithSampleData() error {
	var count int64
	DB.Model(&models.Item{}).Count(&count)
	if count > 0 {
		return nil // Database is not empty, skip populating
	}

	// Sample items
	items := []models.Item{
		{Name: "Laptop", Description: "High-performance laptop", Quantity: 50},
		{Name: "Smartphone", Description: "Latest model smartphone", Quantity: 100},
		{Name: "Tablet", Description: "10-inch tablet", Quantity: 30},
		{Name: "Headphones", Description: "Noise-cancelling headphones", Quantity: 75},
		{Name: "Monitor", Description: "27-inch 4K monitor", Quantity: 25},
	}

	// Sample locations
	locations := []models.Location{
		{Name: "Warehouse A", Capacity: 1000},
		{Name: "Warehouse B", Capacity: 1500},
		{Name: "Store 1", Capacity: 200},
		{Name: "Store 2", Capacity: 250},
		{Name: "Distribution Center", Capacity: 5000},
	}

	// Insert items
	for _, item := range items {
		if err := DB.Create(&item).Error; err != nil {
			return err
		}
	}

	// Insert locations
	for _, location := range locations {
		if err := DB.Create(&location).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetItems(page, limit int, search string) ([]models.Item, int64, error) {
	var items []models.Item
	var total int64

	query := DB.Model(&models.Item{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((page - 1) * limit).Limit(limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func CreateItem(item *models.Item) error {
	return DB.Create(item).Error
}

func UpdateItem(item *models.Item) error {
	return DB.Save(item).Error
}

func DeleteItem(id uint) error {
	return DB.Delete(&models.Item{}, id).Error
}

func GetLocations() ([]models.Location, error) {
	var locations []models.Location
	err := DB.Find(&locations).Error
	return locations, err
}

func CreateLocation(location *models.Location) error {
	return DB.Create(location).Error
}

func UpdateLocation(location *models.Location) error {
	return DB.Save(location).Error
}

func DeleteLocation(id uint) error {
	return DB.Delete(&models.Location{}, id).Error
}

func CreateIssue(issue *models.Issue) error {
	return DB.Create(issue).Error
}

func GetPendingIssues() ([]models.Issue, error) {
	var issues []models.Issue
	err := DB.Where("status = ?", "pending").Find(&issues).Error
	return issues, err
}

func ApproveIssue(id uint) error {
	return DB.Model(&models.Issue{}).Where("id = ?", id).Update("status", "approved").Error
}

func DenyIssue(id uint) error {
	return DB.Model(&models.Issue{}).Where("id = ?", id).Update("status", "denied").Error
}

func GetInventoryReport() ([]models.InventoryReport, error) {
	var report []models.InventoryReport
	err := DB.Model(&models.Item{}).
		Select("id, name, quantity, (SELECT COALESCE(SUM(quantity), 0) FROM issues WHERE item_id = items.id AND status = 'approved') as issued_quantity").
		Find(&report).Error
	return report, err
}

func GetIssueReport() ([]models.Issue, error) {
	var issues []models.Issue
	err := DB.Find(&issues).Error
	return issues, err
}

func GetItemMovementReport(itemID uint) ([]models.Issue, error) {
	var issues []models.Issue
	err := DB.Where("item_id = ?", itemID).Find(&issues).Error
	return issues, err
}
