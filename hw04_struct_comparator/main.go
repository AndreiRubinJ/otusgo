package main

import (
	"fmt"

	"github.com/AndreiRubinJ/otusgo/hw04_struct_comparator/pkg/book"
)

var comparatorNew = book.NewComparator()

func main() {
	var book1 = book.NewBook(1, "Title for book 1", "Author book 1", 2024, 300, 5.0)
	var book2 = book.NewBook(2, "Title for book 2", "Author book 2", 2020, 150, 4.3)
	var book3 = book.NewBook(2, "Title for book 3", "Author book 3", 2018, 800, 4.8)
	bookByYearComparator := book.NewComparatorMode(book.ByYear)
	bookByRateComparator := book.NewComparatorMode(book.ByRate)
	bookSizeComparator := book.NewComparatorMode(book.BySize)
	fmt.Println("0: Test setter should", comparatorNew.SetMode(book.ByYear).Compare(book1, book2))
	fmt.Println("1: should true ", bookByYearComparator.Compare(book1, book2))
	fmt.Println("2: should false ", bookByYearComparator.Compare(book3, book2))
	fmt.Println("3: should true ", bookByRateComparator.Compare(book3, book2))
	fmt.Println("4: should true ", bookSizeComparator.Compare(book1, book3))
}
