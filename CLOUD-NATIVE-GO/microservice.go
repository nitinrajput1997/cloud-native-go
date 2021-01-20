package main

import (
	"google.golang.org/genproto/googleapis/api"
	"golang.org/x/text/message"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("api/echo", api.EchoHandleFunc)
	http.HandleFunc("api/books", api.BooksHandleFunc)
	http.HandleFunc("api/books/", api.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string{
	port := os.Getenv("PORT")
    if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World Cloud Native")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}