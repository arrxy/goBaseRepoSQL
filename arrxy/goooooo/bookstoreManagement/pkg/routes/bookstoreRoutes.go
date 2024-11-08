package routes

import (
	"bookstoreManagement/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBookById).Methods("GET")
}
