package server

import (
	"fmt"
	"template-api-golang/docs"

	xxxxxs "template-api-golang/internal/api/xxxxxs/routes"
	"template-api-golang/internal/healthz"
	"template-api-golang/internal/middlewares"

	"github.com/jindasoft/jinda-platform/xmdw"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title template-api-golang
// @version 1.0.0
// @description template-api-golang
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (s *server) RegisterRoutes() *echo.Echo {
	e := echo.New()
	conf := s.config.App

	// Middleware
	e.Use(
		middleware.Recover(),
		xmdw.LoggerMiddleware(),
		xmdw.TraceIDMiddleware(),
		xmdw.JwtExtractMiddleware(),
		xmdw.ContextMiddleware(conf.Name, conf.Env),
		xmdw.RequestLocaleMiddleware(),
		xmdw.SecureMiddleware(),
	)

	// validator
	e.Validator = xmdw.NewCustomValidator(
		xmdw.Config{Env: conf.Env},
		xmdw.Opt{
			Validations: []xmdw.Validation{
				{Tag: "template", Handler: middlewares.ValidateTemplate}, // for templates
			},
		})

	// swagger
	if conf.Env != "prod" {
		// Set Swagger host from config
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", conf.Port)

		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// health
	healthz.RegisterHealthz(e)

	// API routes
	api := e.Group("/v1")
	xxxxxs.NewRoutes(api, s.mongo)

	return e
}
