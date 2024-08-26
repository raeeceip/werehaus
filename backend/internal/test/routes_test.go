package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// Initialize database connection
	err := database.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Add ping route
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Setup routes
	r.POST("/api/login", handlers.LoginHandler)
	r.POST("/api/logout", handlers.LogoutHandler)
	r.GET("/api/items", handlers.GetItemsHandler)
	r.POST("/api/items", handlers.CreateItemHandler)
	r.PUT("/api/items/:id", handlers.UpdateItemHandler)
	r.DELETE("/api/items/:id", handlers.DeleteItemHandler)
	r.GET("/api/locations", handlers.GetLocationsHandler)
	r.POST("/api/locations", handlers.CreateLocationHandler)
	r.PUT("/api/locations/:id", handlers.UpdateLocationHandler)
	r.DELETE("/api/locations/:id", handlers.DeleteLocationHandler)
	r.POST("/api/issues", handlers.RequestIssueHandler)
	r.GET("/api/issues/pending", handlers.GetPendingIssuesHandler)
	r.POST("/api/issues/:id/approve", handlers.ApproveIssueHandler)
	r.POST("/api/issues/:id/deny", handlers.DenyIssueHandler)
	r.GET("/api/reports/inventory", handlers.InventoryReportHandler)
	r.GET("/api/reports/issues", handlers.IssueReportHandler)
	r.GET("/api/reports/item-movements/:itemId", handlers.ItemMovementReportHandler)

	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestLoginRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	loginData := map[string]string{"username": "testuser", "password": "testpass"}
	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
}

func TestLogoutRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/logout", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
}

func TestGetItemsRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response struct {
		Items []map[string]interface{} `json:"items"`
		Total int                      `json:"total"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestGetLocationsRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/locations", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestGetPendingIssuesRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/issues/pending", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestInventoryReportRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/reports/inventory", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string][]map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "inventory")
}

func TestIssueReportRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/reports/issues", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string][]map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "issues")
}

func TestItemMovementReportRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/reports/item-movements/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string][]map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "movements")
}
