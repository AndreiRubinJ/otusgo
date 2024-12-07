package main

import (
	"fmt"

	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
	"github.com/AndreiRubinJ/otusgo/hw09_serialize/serializer"
)

func main() {
	books := []model.Book{
		{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
		{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
	}
	data, err := serializer.SerializeBooksJSON(books)
	if err != nil {
		fmt.Printf("Serialization failed: %v", err)
	}
	fmt.Println(string(data))

	newBooks, err := serializer.DeserializeBooksJSON(data)
	if err != nil {
		fmt.Printf("Deserialization failed: %v", err)
	}
	for i, book := range newBooks {
		fmt.Printf("Book %d: %v\n", i, book)
		fmt.Printf("Book %d: %v\n", i, book.ID)
	}
	json, err := books[0].MarshalJSON()
	if err != nil {
		return
	}
	fmt.Println(string(json))

	jsonData := `{
		"id": 1,
		"title": "Go Programming",
		"author": "OTUS",
		"year": 2024,
		"size": 256,
		"rate": 5.5
	}`

	var book model.Book
	err = book.UnmarshalJSON([]byte(jsonData))
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Printf("Book: %+v\n", book)

	protoData, err := serializer.SerializeBooksProto(books)
	if err != nil {
		fmt.Printf("Proto Serialization failed: %v\n", err)
		return
	}
	fmt.Println("Proto Data:", protoData)

	protoBooks, err := serializer.DeserializeBooksProto(protoData)
	if err != nil {
		fmt.Printf("Proto Deserialization failed: %v\n", err)
		return
	}
	fmt.Println("Deserialized Proto Books:", protoBooks[0].Author)
}
