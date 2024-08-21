package main

import (
	"api/config"
	"api/pkg/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	r := routes.SetupRouter()
	log.Printf("Server is running on port %s", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, r))
}
