package models

import (
	"github.com/BarTar213/notificator/utils"
	"time"
)

type Notification struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	UserID     int       `json:"user_id"`
	ResourceID int       `json:"resource_id"`
	Tag        string    `json:"tag"`
	CreateDate time.Time `json:"create_date"`
	Read       bool      `json:"read"`
}

func (n *Notification) Reset() {
	n.ID = 0
	n.Message = utils.EmptyStr
	n.UserID = 0
	n.ResourceID = 0
	n.Tag = utils.EmptyStr
	n.CreateDate = time.Time{}
	n.Read = false
}
