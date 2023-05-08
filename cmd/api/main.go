package main

import (
	"github.com/dportaluppi/commerce-integrations-templates/internal/handlers"
	"github.com/dportaluppi/commerce-integrations-templates/internal/mongodb"
	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"github.com/gin-gonic/gin"
)

const (
	mongoURI = "mongodb://localhost:27017"
)

func main() {
	// repo := inmemory.NewRepository()
	repo, err := mongodb.NewRepository(mongoURI)
	if err != nil {
		panic(err)
	}
	service := template.NewService(repo)
	handler := handlers.NewTemplateHandler(service)

	router := gin.Default()

	router.POST("/templates", handler.CreateTemplate)
	router.GET("/templates/:name", handler.GetTemplate)
	router.PUT("/templates/:name", handler.UpdateTemplate)
	router.DELETE("/templates/:name", handler.DeleteTemplate)
	router.POST("/templates/:name/render", handler.RenderTemplate)

	router.Run()
}
