package api

import (
	"log"
	"net/http"

	"github.com/BarTar213/notificator/models"
	"github.com/BarTar213/notificator/storage"
	"github.com/gin-gonic/gin"
)

const (
	templateResource = "template"
)

type TemplateHandlers struct {
	storage storage.Storage
	logger  *log.Logger
}

func NewTemplateHandlers(storage storage.Storage, logger *log.Logger) *TemplateHandlers {
	return &TemplateHandlers{storage: storage, logger: logger}
}

func (h *TemplateHandlers) GetTemplate(c *gin.Context) {

}

func (h *TemplateHandlers) UpdateTemplate(c *gin.Context) {

}

func (h *TemplateHandlers) AddTemplate(c *gin.Context) {
	template := &models.Template{}
	err := c.ShouldBindJSON(template)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Error: invalidRequestBodyErr})
		return
	}



	c.JSON(http.StatusCreated, template)
}

func (h *TemplateHandlers) DeleteTemplate(c *gin.Context) {

}

func (h *TemplateHandlers) ListTemplates(c *gin.Context) {

}
