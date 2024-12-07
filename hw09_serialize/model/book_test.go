package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		input         Book
		expectedValue string
	}{
		{
			Book{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
			`{"id":1,"title":"Book One","author":"Author A","year":2020,"size":300,"rate":4.5}`,
		}, {
			Book{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			`{"id":2,"title":"Book Two","author":"Author B","year":2021,"size":400,"rate":5.0}`,
		},
	}

	for _, test := range tests {
		json, err := test.input.MarshalJSON()
		if err != nil {
			assert.Fail(t, "Error in serialization")
		}

		assert.JSONEq(t, test.expectedValue, string(json), "Serialization err")
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue Book
	}{
		{
			`{"id": 1,"title": "Go Programming","author": "OTUS", "year": 2024, "size": 256, "rate": 5.5}`,
			Book{ID: 1, Title: "Go Programming", Author: "OTUS", Year: 2024, Size: 256, Rate: 5.5},
		},
	}

	for _, test := range tests {
		var book Book
		err := book.UnmarshalJSON([]byte(test.input))
		if err != nil {
			assert.Fail(t, "Error in unmarshalling")
		}

		assert.Equal(t, test.expectedValue, book, "Serialization err")
	}
}
