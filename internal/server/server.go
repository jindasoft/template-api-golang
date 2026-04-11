package server

import (
	"fmt"
	"net/http"
	"template-api-examples/configs"

	"github.com/jindasoft/jinda-platforms/xlogger"
)

type server struct {
	port int
}

func NewServer() *http.Server {
	conf := configs.GetConfig()

	xlogger.Init(
		conf.Server.LogLevel,
		conf.App.Name,
		conf.Server.PrettyPrint,
	)

	newServer := &server{
		port: conf.App.Port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  conf.Server.IdleTimeout,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
	}

	xlogger.SysInfof("Server is running on port: %d", conf.App.Port)

	return server
}
