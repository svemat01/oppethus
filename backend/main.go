package main

import (
	"github.com/svemat01/WebServ/routes"
	"github.com/svemat01/WebServ/server"
	"log"
)

func main() {
	// Server initialization
	app := server.Create()

	// Routes initialization
	err := routes.SetupRoutes(app)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
