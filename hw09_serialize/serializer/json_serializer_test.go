package serializer

import (
	"testing"

	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
	"github.com/stretchr/testify/assert"
)

func TestSerializeBooksJSON(t *testing.T) {
	tests := []struct {
		input         []model.Book
		expectedValue string
	}{
		{
			[]model.Book{
				{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
				{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			},
			`[
				{
					"id": 1,
					"title": "Book One",
					"author": "Author A",
					"year": 2020,
					"size": 300,
					"rate": 4.5
				},
				{
					"id": 2,
					"title": "Book Two",
					"author": "Author B",
					"year": 2021,
					"size": 400,
					"rate": 5.0
				}
			]`,
		},
	}

	for _, test := range tests {
		json, err := SerializeBooksJSON(test.input)
		if err != nil {
			assert.Fail(t, "Error in serialization")
		}

		assert.JSONEq(t, test.expectedValue, string(json), "Serialization err")
	}
}

func TestDeserializeBooksJSON(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue []model.Book
	}{
		{
			`[
				{
					"id": 1,
					"title": "Book One",
					"author": "Author A",
					"year": 2020,
					"size": 300,
					"rate": 4.5
				},
				{
					"id": 2,
					"title": "Book Two",
					"author": "Author B",
					"year": 2021,
					"size": 400,
					"rate": 5.0
				}
			]`,
			[]model.Book{
				{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
				{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			},
		},
	}

	for _, test := range tests {
		data, err := DeserializeBooksJSON([]byte(test.input))
		if err != nil {
			assert.Fail(t, "Error in serialization")
		}

		assert.Equal(t, test.expectedValue, data, "Serialization err")
	}
}
