package router

import (
	"github.com/Gani-laboratory/go-crud/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/todo", controller.AmbilSemuaTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.AmbilTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", controller.TmbhTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", controller.HapusTodo).Methods("DELETE", "OPTIONS")

	return router
}