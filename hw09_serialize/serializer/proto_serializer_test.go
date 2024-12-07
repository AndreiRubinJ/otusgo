package serializer

import (
	"testing"

	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
	"github.com/AndreiRubinJ/otusgo/hw09_serialize/protobuf"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestSerializeBooksProtoJSON(t *testing.T) {
	tests := []struct {
		input         []model.Book
		expectedValue *protobuf.BookList
	}{
		{
			[]model.Book{
				{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
				{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			},
			&protobuf.BookList{
				Books: []*protobuf.Book{
					{Id: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
					{Id: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
				},
			},
		},
	}

	for _, test := range tests {
		data, err := SerializeBooksProto(test.input)
		if err != nil {
			assert.Fail(t, "Error in serialization")
		}
		expectedBytes, _ := proto.Marshal(test.expectedValue)

		assert.Equal(t, expectedBytes, data, "Serialized bytes mismatch")
	}
}

func TestDeserializeBooksProtoJSON(t *testing.T) {
	tests := []struct {
		input         []model.Book
		expectedValue []model.Book
	}{
		{
			[]model.Book{
				{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
				{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			},
			[]model.Book{
				{ID: 1, Title: "Book One", Author: "Author A", Year: 2020, Size: 300, Rate: 4.5},
				{ID: 2, Title: "Book Two", Author: "Author B", Year: 2021, Size: 400, Rate: 5.0},
			},
		},
	}

	for _, test := range tests {
		data, err := SerializeBooksProto(test.input)
		if err != nil {
			assert.Fail(t, "Error in serialization")
		}
		deserializedData, deserializedDatErr := DeserializeBooksProto(data)
		if deserializedDatErr != nil {
			assert.Fail(t, "Error in deserialization")
		}
		assert.Equal(t, test.expectedValue, deserializedData, "Serialization err")
	}
}
