package serializer

import (
	"testing"

	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
	"github.com/AndreiRubinJ/otusgo/hw09_serialize/protobuf"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title  string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Author string  `protobuf:"bytes,3,opt,name=author,proto3" json:"author"`
	Year   int     `protobuf:"varint,4,opt,name=year,proto3" json:"year"`
	Size   int     `protobuf:"varint,5,opt,name=size,proto3" json:"size"`
	Rate   float64 `protobuf:"fixed64,6,opt,name=rate,proto3" json:"rate"`
}

type ProtoBookList struct {
	Books []Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books"`
}

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
