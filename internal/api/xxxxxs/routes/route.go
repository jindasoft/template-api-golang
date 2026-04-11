package routes

import (
	"template-api-examples/internal/api/xxxxxs/handlers"
	"template-api-examples/internal/api/xxxxxs/repositories"
	"template-api-examples/internal/api/xxxxxs/services"

	"github.com/labstack/echo/v5"
)

func NewRoutes(api *echo.Group) {
	repo := repositories.NewXxxxxRepository()
	service := services.NewXxxxxService(repo)

	// xxxxxs
	h := handlers.NewHandler(service)
	g := api.Group("/xxxxxs")
	g.GET("", h.GetXxxxxPaging)
	g.GET("/:id", h.GetXxxxxByID)
	g.POST("", h.PostXxxxx)
}
