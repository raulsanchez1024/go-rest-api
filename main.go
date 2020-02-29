package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID			string		`json:"id"`
	ISBN		string		`json:"isbn"`
	Title		string		`json:"title"`
	Author	*Author		`json:"author"`
}

// Author struct
type Author struct {
	FirstName		string		`json:"firstname"`
	LastName		string		`json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get parameters
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}


func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initiate router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", ISBN: "345678", Title: "Book One", Author: &Author{FirstName: "Raul", LastName: "Sanchez"}})
	books = append(books, Book{ID: "32", ISBN: "987643", Title: "Book Two", Author: &Author{FirstName: "Sofia", LastName: "Sanchez"}})
	books = append(books, Book{ID: "3232", ISBN: "10293", Title: "Book Three", Author: &Author{FirstName: "Raul", LastName: "Sanchez"}})

	//Router handlers/endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Spin up server
	log.Fatal(http.ListenAndServe(":8000", r))
}