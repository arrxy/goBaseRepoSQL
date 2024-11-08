package controller

import (
	"bookstoreManagement/pkg/models"
	_ "bookstoreManagement/pkg/models"
	_ "bookstoreManagement/pkg/utils"
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	"strconv"
	_ "strconv"
)

var NewBook models.Book
var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	newBooks := NewBook.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

var GetBookById = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}
	book, _ := NewBook.GetBookById(int(id))
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&NewBook)
	NewBook.CreateBook()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}
