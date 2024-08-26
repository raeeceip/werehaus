package database

import (
	"go-warehouse-management/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestDatabaseInitialization(t *testing.T) {
	err := InitDB()
	assert.NoError(t, err)

	// Test if the database connection is established
	assert.NotNil(t, DB)
}

func TestPopulateWithSampleData(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Test sample items
	var items []models.Item
	err = DB.Find(&items).Error
	assert.NoError(t, err)
	assert.Len(t, items, 5)
	assert.Equal(t, "Laptop", items[0].Name)
	assert.Equal(t, "Smartphone", items[1].Name)
	assert.Equal(t, "Tablet", items[2].Name)
	assert.Equal(t, "Headphones", items[3].Name)
	assert.Equal(t, "Monitor", items[4].Name)

	// Test sample locations
	var locations []models.Location
	err = DB.Find(&locations).Error
	assert.NoError(t, err)
	assert.Len(t, locations, 5)
	assert.Equal(t, "Warehouse A", locations[0].Name)
	assert.Equal(t, "Warehouse B", locations[1].Name)
	assert.Equal(t, "Store 1", locations[2].Name)
	assert.Equal(t, "Store 2", locations[3].Name)
	assert.Equal(t, "Distribution Center", locations[4].Name)

	// Test admin user creation
	var adminUser models.User
	err = DB.Where("is_admin = ?", true).First(&adminUser).Error
	assert.NoError(t, err)
	assert.Equal(t, "admin", adminUser.Username)
	assert.True(t, adminUser.IsAdmin)
}

func TestCreateAndGetUser(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create a new user
	newUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
		IsAdmin:  false,
	}
	err = CreateUser(newUser)
	assert.NoError(t, err)

	// Get the user by username
	retrievedUser, err := GetUserByUsername("testuser")
	assert.NoError(t, err)
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, "testuser", retrievedUser.Username)
	assert.False(t, retrievedUser.IsAdmin)

	// Check if the password is hashed
	err = bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte("testpassword"))
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create additional users
	users := []models.User{
		{Username: "user1", Password: "password1", IsAdmin: false},
		{Username: "user2", Password: "password2", IsAdmin: false},
	}
	for _, user := range users {
		err = CreateUser(&user)
		assert.NoError(t, err)
	}

	// Get all users
	allUsers, err := GetUsers()
	assert.NoError(t, err)
	assert.Len(t, allUsers, 3) // 2 new users + 1 admin user
}

func TestItemCRUD(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create a new item
	newItem := &models.Item{
		Name:        "New Item",
		Description: "Test item",
		Quantity:    10,
	}
	err = CreateItem(newItem)
	assert.NoError(t, err)
	assert.NotZero(t, newItem.ID)

	// Get items
	items, total, err := GetItems(1, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, int64(6), total) // 5 sample items + 1 new item
	assert.Len(t, items, 6)

	// Update item
	newItem.Quantity = 20
	err = UpdateItem(newItem)
	assert.NoError(t, err)

	// Verify update
	var updatedItem models.Item
	err = DB.First(&updatedItem, newItem.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(20), int64(updatedItem.Quantity))

	// Delete item
	err = DeleteItem(newItem.ID)
	assert.NoError(t, err)

	// Verify deletion
	err = DB.First(&models.Item{}, newItem.ID).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestLocationCRUD(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create a new location
	newLocation := &models.Location{
		Name:     "New Location",
		Capacity: 100,
	}
	err = CreateLocation(newLocation)
	assert.NoError(t, err)
	assert.NotZero(t, newLocation.ID)

	// Get locations
	locations, err := GetLocations()
	assert.NoError(t, err)
	assert.Len(t, locations, 6) // 5 sample locations + 1 new location

	// Update location
	newLocation.Capacity = 200
	err = UpdateLocation(newLocation)
	assert.NoError(t, err)

	// Verify update
	var updatedLocation models.Location
	err = DB.First(&updatedLocation, newLocation.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(200), int64(updatedLocation.Capacity))

	// Delete location
	err = DeleteLocation(newLocation.ID)
	assert.NoError(t, err)

	// Verify deletion
	err = DB.First(&models.Location{}, newLocation.ID).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestIssueCRUD(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create a new issue
	newIssue := &models.Issue{
		ItemID:   1,
		Quantity: 5,
		Status:   "pending",
	}
	err = CreateIssue(newIssue)
	assert.NoError(t, err)
	assert.NotZero(t, newIssue.ID)

	// Get pending issues
	pendingIssues, err := GetPendingIssues()
	assert.NoError(t, err)
	assert.Len(t, pendingIssues, 1)

	// Approve issue
	err = ApproveIssue(newIssue.ID)
	assert.NoError(t, err)

	// Verify approval
	var approvedIssue models.Issue
	err = DB.First(&approvedIssue, newIssue.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "approved", approvedIssue.Status)

	// Create another issue and deny it
	anotherIssue := &models.Issue{
		ItemID:   2,
		Quantity: 3,
		Status:   "pending",
	}
	err = CreateIssue(anotherIssue)
	assert.NoError(t, err)

	err = DenyIssue(anotherIssue.ID)
	assert.NoError(t, err)

	// Verify denial
	var deniedIssue models.Issue
	err = DB.First(&deniedIssue, anotherIssue.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "denied", deniedIssue.Status)
}

func TestReports(t *testing.T) {
	// Reset the database before testing
	err := ResetDatabase()
	assert.NoError(t, err)

	// Create some issues
	issues := []models.Issue{
		{ItemID: 1, Quantity: 5, Status: "approved"},
		{ItemID: 1, Quantity: 3, Status: "pending"},
		{ItemID: 2, Quantity: 2, Status: "approved"},
	}
	for _, issue := range issues {
		err = CreateIssue(&issue)
		assert.NoError(t, err)
	}

	// Test inventory report
	inventoryReport, err := GetInventoryReport()
	assert.NoError(t, err)
	assert.Len(t, inventoryReport, 5) // 5 sample items
	assert.Equal(t, int64(5), int64(inventoryReport[0].IssuedQuantity))
	assert.Equal(t, int64(2), int64(inventoryReport[1].IssuedQuantity))

	// Test issue report
	issueReport, err := GetIssueReport()
	assert.NoError(t, err)
	assert.Len(t, issueReport, 3)

	// Test item movement report
	itemMovements, err := GetItemMovementReport(1)
	assert.NoError(t, err)
	assert.Len(t, itemMovements, 2)
}
