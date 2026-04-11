package server

import (
	"template-api-examples/configs"

	xxxxxs "template-api-examples/internal/api/xxxxxs/routes"
	"template-api-examples/internal/healthz"
	"template-api-examples/internal/middlewares"

	"github.com/jindasoft/jinda-platforms/xmdw"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *server) RegisterRoutes() *echo.Echo {
	conf := configs.GetConfig()

	e := echo.New()

	// Middleware
	e.Use(
		middleware.Recover(),
		xmdw.LoggerMiddleware(),
		xmdw.CorrelationIDMiddleware(),
		xmdw.TraceIDMiddleware(),
		xmdw.SpanIDMiddleware(),
		xmdw.JwtExtractMiddleware(),
		xmdw.ContextMiddleware(conf.App.Name, conf.App.Env),
		xmdw.RequestCultureMiddleware(),
		xmdw.SecureMiddleware(),
	)

	// validator
	e.Validator = xmdw.NewCustomValidator(
		xmdw.Config{Env: conf.App.Env},
		xmdw.Opt{
			Validations: []xmdw.Validation{
				{Tag: "template", Handler: middlewares.ValidateTemplate}, // for templates
			},
		})

	// swagger
	if conf.App.Env != "prod" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// health
	healthz.RegisterHealthz(e)

	// API routes
	api := e.Group("/v1")
	xxxxxs.NewRoutes(api)

	return e
}
