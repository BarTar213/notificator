package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BarTar213/notificator/email"
	"github.com/BarTar213/notificator/models"
	"github.com/BarTar213/notificator/senders"
	"github.com/BarTar213/notificator/storage"
	"github.com/gin-gonic/gin"
)

const (
	templateResource = "template"
	id               = "id"
	name             = "name"

	typeInternal = "internal"
	typeEmail    = "email"
)

type TemplateHandlers struct {
	storage     storage.Storage
	emailClient email.Client
	logger      *log.Logger
}

func NewTemplateHandlers(storage storage.Storage, emailCli email.Client, logger *log.Logger) *TemplateHandlers {
	return &TemplateHandlers{
		storage:     storage,
		emailClient: emailCli,
		logger:      logger,
	}
}

func (h *TemplateHandlers) GetTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	template := &models.Template{ID: id}
	err = h.storage.GetTemplate(template)
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	c.JSON(http.StatusOK, template)
}

func (h *TemplateHandlers) UpdateTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	template := &models.Template{}
	err = c.ShouldBindJSON(template)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Error: invalidRequestBodyErr})
		return
	}

	template.ID = id

	err = h.storage.UpdateTemplate(template)
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	c.JSON(http.StatusOK, template)
}

func (h *TemplateHandlers) AddTemplate(c *gin.Context) {
	template := &models.Template{}
	err := c.ShouldBindJSON(template)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Error: invalidRequestBodyErr})
		return
	}

	_, err = template.Parse(map[string]string{})
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Error: "could not parse template"})
		return
	}

	err = h.storage.AddTemplate(template)
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	c.JSON(http.StatusCreated, template)
}

func (h *TemplateHandlers) DeleteTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidIdParamErr)
		return
	}

	err = h.storage.DeleteTemplate(id)
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *TemplateHandlers) ListTemplates(c *gin.Context) {
	templates, err := h.storage.ListTemplates()
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	c.JSON(http.StatusOK, templates)
}

func (h *TemplateHandlers) SendFromTemplate(c *gin.Context) {
	name := c.Param(name)
	if len(strings.TrimSpace(name)) == 0 {
		c.JSON(http.StatusBadRequest, invalidNameParamErr)
		return
	}

	template := &models.Template{Name: name}
	err := h.storage.GetTemplateByName(template)
	if err != nil {
		handlePostgresError(c, h.logger, err, templateResource)
		return
	}

	notificationType := c.Query("type")
	if len(notificationType) == 0 || (notificationType != typeInternal && notificationType != typeEmail) {
		c.JSON(http.StatusBadRequest, &models.Response{Error: "Query value 'type' should be provided. Possible values are: internal, email"})
		return
	}

	switch notificationType {
	case typeInternal:
		sender := &senders.Internal{}
		err := c.ShouldBindJSON(sender)
		if err != nil {
			c.JSON(http.StatusBadRequest, &models.Response{Error: invalidRequestBodyErr})
			return
		}
		go h.SendInternal(sender, template)
	case typeEmail:
		sender := &senders.Email{}
		err := c.ShouldBindJSON(sender)
		if err != nil {
			c.JSON(http.StatusBadRequest, &models.Response{Error: invalidRequestBodyErr})
			return
		}
		go h.SendEmail(sender, template)
	}

	c.JSON(http.StatusAccepted, struct{}{})
}

func (h *TemplateHandlers) SendInternal(sender *senders.Internal, template *models.Template) {
	message, err := template.Parse(sender.Data)
	if err != nil {
		h.logger.Printf("template parse: %s", err)
		return
	}

	err = sender.Send(h.storage, message)
	if err != nil {
		h.logger.Printf("internal sending: %s", err)
		return
	}
}

func (h *TemplateHandlers) SendEmail(sender *senders.Email, template *models.Template) {
	message, err := template.Parse(sender.Data)
	if err != nil {
		h.logger.Printf("template parse: %s", err)
		return
	}

	err = sender.Send(h.emailClient, template.Title, message, template.HTML)
	if err != nil {
		h.logger.Printf("email sending: %s", err)
		return
	}
}
