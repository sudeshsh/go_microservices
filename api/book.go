package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Name Author and ISBN
type Book struct {
	//define the Book
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// Books slice of all known books
var Books = []Book{
	Book{Title: "Go Microservice", Author: "S. M. Shinde", ISBN: "0123456789"},
	Book{Title: "Cloud", Author: "A. B. Lol", ISBN: "0987654321"},
}

// BooksHandleFunc tu used as http.handleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content_type", "aplication/json; charset=utf-8")
	w.Write(b)
}
