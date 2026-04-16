package main

import (
	"log"
	"template-api-golang/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
