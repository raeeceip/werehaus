package main

import (
    "log"
    "net/http"

    "go-warehouse-management/internal/database"
    "go-warehouse-management/internal/handlers"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    db, err := database.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/", handlers.HomeHandler)
    r.Get("/items", handlers.ItemsHandler)
    r.Get("/locations", handlers.LocationsHandler)
    r.Get("/issue", handlers.IssueHandler)
    r.Get("/reports", handlers.ReportsHandler)

    r.Post("/api/items", handlers.CreateItemHandler)
    r.Post("/api/locations", handlers.CreateLocationHandler)
    r.Post("/api/issue", handlers.IssueItemHandler)

    fileServer := http.FileServer(http.Dir("./web/static"))
    r.Handle("/static/*", http.StripPrefix("/static", fileServer))

    log.Println("Server starting on :3000")
    log.Fatal(http.ListenAndServe(":3000", r))
}
