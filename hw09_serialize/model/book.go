package model

import (
	"encoding/json"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (book *Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(book),
	})
}

func (book *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	err := json.Unmarshal(data, &struct {
		*Alias
	}{
		Alias: (*Alias)(book),
	})
	return err
}
