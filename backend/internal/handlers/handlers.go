package handlers

import (
	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/models"
	"html/template"
	"net/http"
	"strconv"
)

type Item struct {
	Name        string
	Description string
	Quantity    int
}

type Location struct {
	Name        string
	Description string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/home.html"))
	tmpl.Execute(w, nil)
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/items.html"))
	tmpl.Execute(w, nil)
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/locations.html"))
	tmpl.Execute(w, nil)
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/issue.html"))
	tmpl.Execute(w, nil)
}

func ReportsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/reports.html"))
	tmpl.Execute(w, nil)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Extract item details from the form
	name := r.FormValue("name")
	description := r.FormValue("description")
	quantity := r.FormValue("quantity")

	// Validate input
	if name == "" || quantity == "" {
		http.Error(w, "Name and quantity are required", http.StatusBadRequest)
		return
	}

	// Convert quantity to integer
	qty, err := strconv.Atoi(quantity)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Create a new item (you'll need to implement this function in your data layer)
	newItem := &models.Item{
		Name:        name,
		Description: description,
		Quantity:    qty,
	}
	err = database.CreateItem(newItem)
	if err != nil {
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	// Redirect to the items page after successful creation
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}

func CreateLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	newLocation := &models.Location{
		Name: name,
	}

	err = database.CreateLocation(newLocation)
	if err != nil {
		http.Error(w, "Error creating location", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/locations", http.StatusSeeOther)
}

func IssueItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	itemID := r.FormValue("item_id")
	locationID := r.FormValue("location_id")
	quantity := r.FormValue("quantity")

	if itemID == "" || locationID == "" || quantity == "" {
		http.Error(w, "Item ID, Location ID, and quantity are required", http.StatusBadRequest)
		return
	}

	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	locationIDInt, err := strconv.Atoi(locationID)
	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	err = database.IssueItem(itemIDInt, locationIDInt, quantityInt)
	if err != nil {
		http.Error(w, "Error issuing item", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
