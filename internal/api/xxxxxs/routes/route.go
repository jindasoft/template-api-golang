package routes

import (
	"template-api-golang/internal/api/xxxxxs/handlers"
	"template-api-golang/internal/api/xxxxxs/repositories"
	"template-api-golang/internal/api/xxxxxs/services"

	"github.com/jindasoft/jinda-platform/xdb"
	"github.com/labstack/echo/v5"
)

func NewRoutes(api *echo.Group, mongo xdb.MongoService) {
	repo := repositories.NewXxxxxRepository(mongo)
	service := services.NewXxxxxService(repo)

	// xxxxxs
	h := handlers.NewHandler(service)
	g := api.Group("/xxxxxs")
	g.GET("", h.GetXxxxxPaging)
	g.GET("/:id", h.GetXxxxxByID)
	g.POST("", h.PostXxxxx)
	g.PUT("/:id", h.PutXxxxx)
	g.PUT("/:id/status", h.PutSetActiveXxxxx)
	g.DELETE("/:id", h.DeleteXxxxx)
}
