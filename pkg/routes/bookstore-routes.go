package routes

import (
	"net/http"

	logger "github.com/erfan-motallebi/bookstore/middleware"
	"github.com/erfan-motallebi/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookstoreRoutes(router *mux.Router) {
	router.Handle("/books", logger.LoggerMiddleware(http.HandlerFunc(controllers.GetBooks))).Methods("GET")
	router.Handle("/books", logger.LoggerMiddleware(http.HandlerFunc(controllers.CreateBook))).Methods("POST")
	router.Handle("/books/{bookId}", logger.LoggerMiddleware(http.HandlerFunc(controllers.FindBookById))).Methods("GET")
	router.Handle("/books/{bookId}", logger.LoggerMiddleware(http.HandlerFunc(controllers.UpdateBook))).Methods("PUT")
	router.Handle("/books/{bookId}", logger.LoggerMiddleware(http.HandlerFunc(controllers.DeleteBook))).Methods("DELETE")
}
