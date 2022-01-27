package handlers

import (
	"book_api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		log.Fatal("Error in decode json")
		writer.WriteHeader(400)
		msg := models.ErrorMessage{Message: "Error in decode json"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var newBookId = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(models.ErrorMessage{Message: "Not valid if of book"})
		return
	}
	book, ok := models.FindBookById(id)
	if !ok {
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(models.ErrorMessage{Message: "Not found book"})
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(models.ErrorMessage{Message: "Not valid if of book"})
		return
	}
	book, ok := models.FindBookById(id)
	if !ok {
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(models.ErrorMessage{Message: "Not found book"})
		return
	}
	var updateBook models.Book
	err = json.NewDecoder(request.Body).Decode(&updateBook)
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(models.ErrorMessage{Message: "Not valid request body for update book"})
		return
	}
	book.Name = updateBook.Name
	book.Description = updateBook.Description
	book.PublishedYear = updateBook.PublishedYear
	book.Author = updateBook.Author
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(book)

}
