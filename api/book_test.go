package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Automating Junos with Ansible", Author: "Sean Sawtell", ISBN: "0123456"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Automating Junos with Ansible","author":"Sean Sawtell","isbn":"0123456"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Automating Junos with Ansible","author":"Sean Sawtell","isbn":"0123456"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Automating Junos with Ansible", Author: "Sean Sawtell", ISBN: "0123456"},
		book, "Book JSON unmarshalling wrong.")
}
