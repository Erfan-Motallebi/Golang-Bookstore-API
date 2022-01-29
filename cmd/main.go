package main

import (
	"log"
	"net/http"

	logger "github.com/erfan-motallebi/bookstore/middleware"
	"github.com/erfan-motallebi/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	routes.RegisterBookstoreRoutes(app)

	app.Handle("/", logger.LoggerMiddleware(app))

	log.Fatal(http.ListenAndServe("localhost:8050", app), nil)
}
