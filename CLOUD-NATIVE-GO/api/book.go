package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//book type with name , author and isbn
type Book struct {
	//define the book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// Books slice of all known  books
var Books = []Book{
	Book{Title: "Engineering Mathematics", Author: "S.Chand", ISBN: "0345678921"},
	Book{Title: "Cloud Native Go", Author: "Nitin Rajput", ISBN: "0000000000"}
}

//BookHandleFunc to be used as http.Handlefunc for Book-Api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	//implement logic for /api/books
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

//BookHandleFunc to be used as http.Handlefunc for Book-Api
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	//implement logic for /api/book/<isbn>
	isbn := r.URL.Path[len("/api/books"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

//AllBooks returns slice of all books
func AllBooks() []Book {

}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {

}

//ToJSON to be used for marshalling of book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJSON to be used for unmarshalling of book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

//BooksHnadleFunc to be used as http.HandleFunc for book api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
