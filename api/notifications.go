package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BarTar213/notificator/models"
	"github.com/BarTar213/notificator/storage"
	"github.com/gin-gonic/gin"
)

const (
	notificationResource = "notification"
)

type NotificationHandlers struct {
	storage storage.Storage
	logger  *log.Logger
}

func NewNotificationHandlers(storage storage.Storage, logger *log.Logger) *NotificationHandlers {
	return &NotificationHandlers{storage: storage, logger: logger}
}

func (h *NotificationHandlers) GetNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	notification := &models.Notification{ID: id}
	err = h.storage.GetNotification(notification)
	if err != nil {
		handlePostgresError(c, h.logger, err, notificationResource)
		return
	}

	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandlers) UpdateNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	read, err := strconv.ParseBool(c.Query("read"))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidReadParamErr)
		return
	}

	err = h.storage.ReadNotification(id, read)
	if err != nil {
		handlePostgresError(c, h.logger, err, notificationResource)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *NotificationHandlers) DeleteNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	err = h.storage.DeleteNotification(id)
	if err != nil {
		handlePostgresError(c, h.logger, err, notificationResource)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *NotificationHandlers) ListNotifications(c *gin.Context) {
	notifications, err := h.storage.ListNotifications()
	if err != nil {
		handlePostgresError(c, h.logger, err, notificationResource)
		return
	}

	c.JSON(http.StatusOK, notifications)
}
