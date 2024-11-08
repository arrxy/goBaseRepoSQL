package models

import (
	"bookstoreManagement/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook(book *Book) *Book {
	db.Save(&book)
	return book
}

func (b *Book) DeleteBook(id int) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}
