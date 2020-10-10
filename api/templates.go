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
	templateResource = "template"
	id               = "id"
)

type TemplateHandlers struct {
	storage storage.Storage
	logger  *log.Logger
}

func NewTemplateHandlers(storage storage.Storage, logger *log.Logger) *TemplateHandlers {
	return &TemplateHandlers{storage: storage, logger: logger}
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

}
