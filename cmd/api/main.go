package main

import (
	"github.com/dportaluppi/commerce-integrations-templates/internal/handlers"
	"github.com/dportaluppi/commerce-integrations-templates/internal/inmemory"
	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := inmemory.NewRepository()
	service := template.NewService(repo)
	handler := handlers.NewTemplateHandler(service)

	router := gin.Default()

	router.POST("/templates", handler.CreateTemplate)
	router.GET("/templates/:name", handler.GetTemplate)
	router.POST("/templates/:name/render", handler.RenderTemplate)

	router.Run()
}
