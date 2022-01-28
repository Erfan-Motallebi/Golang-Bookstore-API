package routes

import (
	"github.com/erfan-motallebi/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookstoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.FindBookById).Methods("GET")
}