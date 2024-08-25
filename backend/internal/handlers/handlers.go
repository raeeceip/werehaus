package handlers

import (
	"net/http"
	"strconv"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/models"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// TODO: Implement login logic
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func LogoutHandler(c *gin.Context) {
	// TODO: Implement logout logic
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func GetItemsHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	items, total, err := database.GetItems(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items, "total": total})
}

func CreateItemHandler(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func UpdateItemHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)
	if err := database.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func DeleteItemHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting item"})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetLocationsHandler(c *gin.Context) {
	locations, err := database.GetLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func CreateLocationHandler(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}

func UpdateLocationHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location.ID = uint(id)
	if err := database.UpdateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating location"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func DeleteLocationHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DeleteLocation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting location"})
		return
	}

	c.Status(http.StatusNoContent)
}

func RequestIssueHandler(c *gin.Context) {
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateIssue(&issue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating issue"})
		return
	}

	c.JSON(http.StatusCreated, issue)
}

func GetPendingIssuesHandler(c *gin.Context) {
	issues, err := database.GetPendingIssues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching pending issues"})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func ApproveIssueHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.ApproveIssue(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error approving issue"})
		return
	}

	c.Status(http.StatusOK)
}

func DenyIssueHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DenyIssue(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error denying issue"})
		return
	}

	c.Status(http.StatusOK)
}

func InventoryReportHandler(c *gin.Context) {
	report, err := database.GetInventoryReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating inventory report"})
		return
	}

	c.JSON(http.StatusOK, report)
}

func IssueReportHandler(c *gin.Context) {
	report, err := database.GetIssueReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating issue report"})
		return
	}

	c.JSON(http.StatusOK, report)
}

func ItemMovementReportHandler(c *gin.Context) {
	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	report, err := database.GetItemMovementReport(uint(itemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating item movement report"})
		return
	}

	c.JSON(http.StatusOK, report)
}
