package main

import (
	"log"
	"net/http"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"}, // Update this with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	// Routes
	r.Post("/api/login", handlers.LoginHandler)
	r.Post("/api/logout", handlers.LogoutHandler)

	r.Route("/api/items", func(r chi.Router) {
		r.Get("/", handlers.GetItemsHandler)
		r.Post("/", handlers.CreateItemHandler)
		r.Put("/{id}", handlers.UpdateItemHandler)
		r.Delete("/{id}", handlers.DeleteItemHandler)
	})

	r.Route("/api/locations", func(r chi.Router) {
		r.Get("/", handlers.GetLocationsHandler)
		r.Post("/", handlers.CreateLocationHandler)
		r.Put("/{id}", handlers.UpdateLocationHandler)
		r.Delete("/{id}", handlers.DeleteLocationHandler)
	})

	r.Route("/api/issues", func(r chi.Router) {
		r.Post("/", handlers.RequestIssueHandler)
		r.Get("/pending", handlers.GetPendingIssuesHandler)
		r.Post("/{id}/approve", handlers.ApproveIssueHandler)
		r.Post("/{id}/deny", handlers.DenyIssueHandler)
	})

	r.Route("/api/reports", func(r chi.Router) {
		r.Get("/inventory", handlers.InventoryReportHandler)
		r.Get("/issues", handlers.IssueReportHandler)
		r.Get("/item-movements/{itemId}", handlers.ItemMovementReportHandler)
	})

	log.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
