package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/erfan-motallebi/bookstore/pkg/models"
	"github.com/erfan-motallebi/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	books := models.GetAllBooks()

	// BooksJson, err := json.Marshal(books)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
	// io.WriteString(w, string(BooksJson))

	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	bJson, err := json.Marshal(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	io.WriteString(w, string(bJson))
}

func FindBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ID := mux.Vars(r)["bookId"]
	IntOfID, _ := strconv.Atoi(ID)
	book := models.FindBookById(IntOfID)
	json.NewEncoder(w).Encode(book)

}
