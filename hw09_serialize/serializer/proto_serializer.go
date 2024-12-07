package serializer

import (
	"github.com/AndreiRubinJ/otusgo/hw09_serialize/model"
	"github.com/AndreiRubinJ/otusgo/hw09_serialize/protobuf"
	"google.golang.org/protobuf/proto"
)

func SerializeBooksProto(books []model.Book) ([]byte, error) {
	pbBooks := make([]*protobuf.Book, len(books))
	for i, book := range books {
		pbBooks[i] = &protobuf.Book{
			Id:     int32(book.ID),
			Title:  book.Title,
			Author: book.Author,
			Year:   int32(book.Year),
			Size:   int32(book.Size),
			Rate:   book.Rate,
		}
	}
	return proto.Marshal(&protobuf.BookList{Books: pbBooks})
}

func DeserializeBooksProto(data []byte) ([]model.Book, error) {
	var pbBookList protobuf.BookList
	err := proto.Unmarshal(data, &pbBookList)
	if err != nil {
		return nil, err
	}

	books := make([]model.Book, len(pbBookList.GetBooks()))
	for i, pbBook := range pbBookList.Books {
		books[i] = model.Book{
			ID:     int(pbBook.Id),
			Title:  pbBook.Title,
			Author: pbBook.Author,
			Year:   int(pbBook.Year),
			Size:   int(pbBook.Size),
			Rate:   pbBook.Rate,
		}
	}
	return books, nil
}
