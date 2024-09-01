package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBook(id int, title, author string, year, size int, rate float32) *Book {
	return &Book{id, title, author, year, size, rate}
}

func (b *Book) SetID(id int) *Book {
	b.id = id
	return b
}
func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetTitle(title string) *Book {
	b.title = title
	return b
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) *Book {
	b.author = author
	return b
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) *Book {
	b.year = year
	return b
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) *Book {
	b.size = size
	return b
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float32) *Book {
	b.rate = rate
	return b
}

func (b *Book) Rate() float32 {
	return b.rate
}
