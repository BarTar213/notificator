package models

import "time"

type Notification struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	UserID     int       `json:"user_id"`
	ResourceID int       `json:"resource_id"`
	Tag        string    `json:"tag"`
	CreateDate time.Time `json:"create_date"`
	Read       bool      `json:"read"`
}
