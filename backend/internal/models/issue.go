package models

import "time"

type Issue struct {
    ID         int       `json:"id"`
    ItemID     int       `json:"item_id"`
    LocationID int       `json:"location_id"`
    Quantity   int       `json:"quantity"`
    IssueDate  time.Time `json:"issue_date"`
}
