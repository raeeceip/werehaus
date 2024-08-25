package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-warehouse-management/internal/database"
	"go-warehouse-management/internal/models"

	"github.com/go-chi/chi/v5"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement login logic
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logout logic
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")

	items, err := database.GetItems(page, limit, search)
	if err != nil {
		http.Error(w, "Error fetching items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.CreateItem(&item)
	if err != nil {
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item.ID = id
	err = database.UpdateItem(&item)
	if err != nil {
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := database.DeleteItem(id)
	if err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := database.GetLocations()
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

func CreateLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location models.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.CreateLocation(&location)
	if err != nil {
		http.Error(w, "Error creating location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(location)
}

func UpdateLocationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var location models.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	location.ID = id
	err = database.UpdateLocation(&location)
	if err != nil {
		http.Error(w, "Error updating location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func DeleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := database.DeleteLocation(id)
	if err != nil {
		http.Error(w, "Error deleting location", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RequestIssueHandler(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue
	err := json.NewDecoder(r.Body).Decode(&issue)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = database.CreateIssue(&issue)
	if err != nil {
		http.Error(w, "Error creating issue", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(issue)
}

func GetPendingIssuesHandler(w http.ResponseWriter, r *http.Request) {
	issues, err := database.GetPendingIssues()
	if err != nil {
		http.Error(w, "Error fetching pending issues", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}

func ApproveIssueHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := database.ApproveIssue(id)
	if err != nil {
		http.Error(w, "Error approving issue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DenyIssueHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := database.DenyIssue(id)
	if err != nil {
		http.Error(w, "Error denying issue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func InventoryReportHandler(w http.ResponseWriter, r *http.Request) {
	report, err := database.GetInventoryReport()
	if err != nil {
		http.Error(w, "Error generating inventory report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func IssueReportHandler(w http.ResponseWriter, r *http.Request) {
	report, err := database.GetIssueReport()
	if err != nil {
		http.Error(w, "Error generating issue report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func ItemMovementReportHandler(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemId")
	report, err := database.GetItemMovementReport(itemID)
	if err != nil {
		http.Error(w, "Error generating item movement report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
