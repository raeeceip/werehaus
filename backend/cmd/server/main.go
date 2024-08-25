package main

import (
	"log"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"} // Update this with your frontend URL
	r.Use(cors.New(config))

	// Routes
	api := r.Group("/api")
	{
		api.POST("/login", handlers.LoginHandler)
		api.POST("/logout", handlers.LogoutHandler)

		items := api.Group("/items")
		{
			items.GET("/", handlers.GetItemsHandler)
			items.POST("/", handlers.CreateItemHandler)
			items.PUT("/:id", handlers.UpdateItemHandler)
			items.DELETE("/:id", handlers.DeleteItemHandler)
		}

		locations := api.Group("/locations")
		{
			locations.GET("/", handlers.GetLocationsHandler)
			locations.POST("/", handlers.CreateLocationHandler)
			locations.PUT("/:id", handlers.UpdateLocationHandler)
			locations.DELETE("/:id", handlers.DeleteLocationHandler)
		}

		issues := api.Group("/issues")
		{
			issues.POST("/", handlers.RequestIssueHandler)
			issues.GET("/pending", handlers.GetPendingIssuesHandler)
			issues.POST("/:id/approve", handlers.ApproveIssueHandler)
			issues.POST("/:id/deny", handlers.DenyIssueHandler)
		}

		reports := api.Group("/reports")
		{
			reports.GET("/inventory", handlers.InventoryReportHandler)
			reports.GET("/issues", handlers.IssueReportHandler)
			reports.GET("/item-movements/:itemId", handlers.ItemMovementReportHandler)
		}
	}

	log.Println("Server starting on :3000")
	r.Run(":3000")
}
