package serializer

import (
	"encoding/json"

	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
)

func SerializeBooksJSON(books []model.Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksJSON(data []byte) ([]model.Book, error) {
	var books []model.Book
	err := json.Unmarshal(data, &books)
	return books, err
}
