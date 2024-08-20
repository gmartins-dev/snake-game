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
    log.Fatal(http.ListenAndServe(config.ServerAddress, r))
}
