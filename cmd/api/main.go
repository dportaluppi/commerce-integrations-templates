package main

import (
	"github.com/dportaluppi/commerce-integrations-templates/internal/config"
	"github.com/dportaluppi/commerce-integrations-templates/internal/handlers"
	"github.com/dportaluppi/commerce-integrations-templates/internal/mongodb"
	"github.com/dportaluppi/commerce-integrations-templates/pkg/template"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// repo := inmemory.NewRepository()
	repo, err := mongodb.NewRepository(cfg.MongoURI)
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
