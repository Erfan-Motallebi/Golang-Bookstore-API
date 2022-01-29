package main

import (
	"log"
	"net/http"

	"github.com/erfan-motallebi/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	routes.RegisterBookstoreRoutes(app)

	app.Handle("/", app)

	log.Fatal(http.ListenAndServe("localhost:8050", app), nil)
}
