package api

import (
	"encoding/json"
	"net/http"
)

// Book type with name, author and isbn
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON to be used for marshalling of Book Type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
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
	Book{Title: "JSNAPy Day One", Author: "Premesh Shah", ISBN: "0000001"},
	Book{Title: "Automating Junos with Ansible", Author: "Sean Sawtell", ISBN: "0123456"}}

// BooksHandleFunc to be used as httpHandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
