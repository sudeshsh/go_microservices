package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Go Microservice", Author: "S. M. Shinde", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Go Microservice","author":"S. M. Shinde","isbn":"0123456789"}`,
		string(json), "Book json marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Go Microservice","author":"S. M. Shinde","isbn":"0123456789"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Go Microservice", Author: "S. M. Shinde", ISBN: "0123456789"},
		book, "Book json unmarshalling wrong")
}
