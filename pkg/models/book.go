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

func UpdateBookById(b Book, ID int) Book {
	var book Book
	db.Model(&Book{}).Where("id = ?", ID).Updates(Book{Name: b.Name, Author: b.Author, Nickname: b.Nickname})
	db.Find(&book, "id = ? ", ID)
	return book
}

func (b *Book) DeleteBookById(ID int) map[string]interface{} {
	delMessage := map[string]interface{}{

		"message": "Book was deleted successfully",
		"success": true,
	}
	db.Delete(&Book{}, ID)

	return delMessage
}
