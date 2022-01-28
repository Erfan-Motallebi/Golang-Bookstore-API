package models

import (
	"github.com/erfan-motallebi/bookstore/pkg/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Nickname string `json:"nickname"`
}

var db *gorm.DB

func init() {
	database.ConnectDB()
	db = database.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func FindBookById(ID int) Book {
	var book Book
	db.Find(&book, "id = ?", ID)
	return book
}
