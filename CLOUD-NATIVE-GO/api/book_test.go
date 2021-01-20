package api

import (
	"encoding/json"
	"testing"
	
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Nitin Rajput", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go", "author": "Nitin Rajput", "isbn": "0123456789"}`,
       string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go", "author": "Nitin Rajput", "isbn": "0123456789"}`)
	book := FromJSON(json)
	
	assert.Equal(t, Book{Title:"Cloud Native Go", Author: "Nitin Rajput", ISBN: "0123456789"},
       book, "Book JSON unmarshalling wrong.")
}