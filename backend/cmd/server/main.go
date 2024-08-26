package main

import (
	"context"
	"log"
	"os"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/handlers"

	apitoolkit "github.com/apitoolkit/apitoolkit-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	apitoolkitCfg := apitoolkit.Config{
		APIKey: os.Getenv("APITOOLKIT_API_KEY"),
	}

	// Check if the API key is set
	if apitoolkitCfg.APIKey == "" {
		log.Fatal("APITOOLKIT_API_KEY environment variable is not set")
	}
	// Initialize the client using your apitoolkit generated apikey
	apitoolkitClient, err := apitoolkit.NewClient(context.Background(), apitoolkitCfg)
	if err != nil {
		log.Fatalf("Failed to initialize APIToolkit client: %v", err)
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Add APIToolkit middleware
	r.Use(apitoolkit.GinMiddleware(apitoolkitClient))

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"}
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
