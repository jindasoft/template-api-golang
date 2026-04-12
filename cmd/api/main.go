package main

import (
	"log"
	_ "template-api-golang/docs"
	"template-api-golang/internal/server"
)

// @title Template API
// @version 1.0.0
// @description Template API Examples
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
