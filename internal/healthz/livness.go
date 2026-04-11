package healthz

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type livenessStats struct {
	Success bool   `json:"success" example:"true"`
	Type    string `json:"type" example:"liveness"`
	Message string `json:"message" example:"Liveness healthy"`
}

// Liveness
// @Summary      Get liveness health status
// @Description  Checks if the application is alive and running (heartbeat check)
// @Tags         health
// @Produce      json
// @Success      200  {object}  livenessStats
// @Router       /liveness [get]
func liveness(c *echo.Context) error {
	liveness := livenessStats{
		Success: true,
		Type:    "liveness",
		Message: "Liveness healthy",
	}

	return c.JSON(http.StatusOK, liveness)
}
