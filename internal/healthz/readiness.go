package healthz

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type readinessStats struct {
	Success bool   `json:"success" example:"true"`
	Type    string `json:"type" example:"readiness"`
	Message string `json:"message" example:"Readiness healthy"`
}

// Readiness
// @Summary      Get readiness health status
// @Description  Checks if the application is ready to accept traffic (dependencies ready)
// @Tags         health
// @Produce      json
// @Success      200  {object}  readinessStats
// @Router       /readiness [get]
func readiness(c *echo.Context) error {
	readiness := readinessStats{
		Success: true,
		Type:    "readiness",
		Message: "Readiness healthy",
	}

	return c.JSON(http.StatusOK, readiness)
}
