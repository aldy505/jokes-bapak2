package main

import (
	"log"
	"net/http"

	"github.com/aldy505/jokes-bapak2-api/api/routes"
)

func main() {
	routes := routes.Setup()
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: routes,
	}
	log.Printf("[info] Server is running on http://localhost:3000")
	server.ListenAndServe()
}
