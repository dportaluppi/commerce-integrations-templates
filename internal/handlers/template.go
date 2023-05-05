package handlers

import (
	"net/http"

	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	service template.Service
}

func NewTemplateHandler(service template.Service) *TemplateHandler {
	return &TemplateHandler{service: service}
}

func (h *TemplateHandler) CreateTemplate(c *gin.Context) {
	var input struct {
		Name    string         `json:"name"`
		Content map[string]any `json:"content"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmpl, err := h.service.Create(input.Name, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tmpl)
}

func (h *TemplateHandler) GetTemplate(c *gin.Context) {
	name := c.Param("name")

	tmpl, err := h.service.Get(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": tmpl.Name, "content": tmpl.Content})
}

func (h *TemplateHandler) RenderTemplate(c *gin.Context) {
	name := c.Param("name")

	var input map[string]interface{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmpl, err := h.service.Get(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	rendered, err := tmpl.Render(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rendered)
}
