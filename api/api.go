package api

import (
	"github.com/BarTar213/notificator/models"
	"log"
	"net/http"
	"sync"

	"github.com/BarTar213/notificator/config"
	"github.com/BarTar213/notificator/email"
	"github.com/BarTar213/notificator/storage"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Port             string
	Router           *gin.Engine
	Config           *config.Config
	Storage          storage.Storage
	EmailClient      email.Client
	NotificationPool *sync.Pool
	TemplatePool     *sync.Pool
	Logger           *log.Logger
}

func WithConfig(conf *config.Config) func(a *Api) {
	return func(a *Api) {
		a.Config = conf
	}
}

func WithLogger(logger *log.Logger) func(a *Api) {
	return func(a *Api) {
		a.Logger = logger
	}
}

func WithStorage(storage storage.Storage) func(a *Api) {
	return func(a *Api) {
		a.Storage = storage
	}
}

func WithEmailClient(emailCli email.Client) func(a *Api) {
	return func(a *Api) {
		a.EmailClient = emailCli
	}
}

func WithNotificationPool() func(a *Api) {
	return func(a *Api) {
		a.NotificationPool = &sync.Pool{
			New: func() interface{} {
				return new(models.Notification)
			},
		}
	}
}

func WithTemplatePool() func(a *Api) {
	return func(a *Api) {
		a.TemplatePool = &sync.Pool{
			New: func() interface{} {
				return new(models.Template)
			},
		}
	}
}

func NewApi(options ...func(api *Api)) *Api {
	a := &Api{
		Router: gin.Default(),
	}
	a.Router.Use(gin.Recovery())

	for _, option := range options {
		option(a)
	}

	a.Router.GET("/", a.health)

	th := NewTemplateHandlers(a.Storage, a.EmailClient, a.TemplatePool, a.Logger)
	nh := NewNotificationHandlers(a.Storage, a.NotificationPool, a.Logger)

	templates := a.Router.Group("/templates")
	{
		templates.GET("/:id", th.GetTemplate)
		templates.GET("", th.ListTemplates)
		templates.PUT("/:id", th.UpdateTemplate)
		templates.POST("", th.AddTemplate)
		templates.DELETE("/:id", th.DeleteTemplate)

		templates.POST("/:name/send", th.SendFromTemplate)
	}

	notifications := a.Router.Group("/notifications")
	{
		notifications.GET("/:id", nh.GetNotification)
		notifications.GET("", nh.ListNotifications)
		notifications.PATCH("/:id", nh.UpdateNotification)
		notifications.DELETE("/:id", nh.DeleteNotification)
	}

	return a
}

func (a *Api) Run() error {
	return a.Router.Run(a.Config.Api.Port)
}

func (a *Api) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "healthy")
}
