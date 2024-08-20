package routes

import (
	"api/pkg/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/start", handlers.StartGame).Methods("POST")
    r.HandleFunc("/validate", handlers.ValidateMoves).Methods("POST")
    return r
}
