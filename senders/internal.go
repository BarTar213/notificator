package senders

import (
	"time"

	"github.com/BarTar213/notificator/models"
	"github.com/BarTar213/notificator/storage"
)

type Internal struct {
	ResourceID int               `json:"resource_id"`
	Resource   string            `json:"resource"`
	Tag        string            `json:"tag"`
	Recipients []int             `json:"recipients"`
	Data       map[string]string `json:"data"`
}

func (i *Internal) Send(s storage.Storage, message string) error {
	createDate := time.Now()

	notifications := make([]*models.Notification, len(i.Recipients))
	for j := 0; j < len(i.Recipients); j++ {
		notifications[j] = &models.Notification{
			Message:    message,
			UserID:     i.Recipients[j],
			Resource:   i.Resource,
			ResourceID: i.ResourceID,
			Tag:        i.Tag,
			CreateDate: createDate,
		}
	}

	return s.BatchAddNotifications(notifications)
}

func (i *Internal) GetData() map[string]string {
	return i.Data
}
