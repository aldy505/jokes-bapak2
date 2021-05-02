package main

import (
	"net/http"

	"github.com/aldy505/jokes-bapak2-api/api/routes"
)

func main() {
	routes := routes.Setup()
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: routes,
	}
	server.ListenAndServe()
}
