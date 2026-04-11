package healthz

import (
	"github.com/labstack/echo/v5"
)

func RegisterHealthz(e *echo.Echo) {
	e.GET("/liveness", liveness)
	e.GET("/readiness", readiness)
}
