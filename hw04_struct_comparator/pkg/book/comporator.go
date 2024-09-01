package book

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type Comparator struct {
	mode CompareMode
}

func NewComparatorMode(mode CompareMode) *Comparator {
	return &Comparator{mode}
}
func NewComparator() *Comparator {
	return &Comparator{}
}
func (compare *Comparator) SetMode(mode CompareMode) *Comparator {
	compare.mode = mode
	return compare
}

func (compare *Comparator) Compare(book1, book2 *Book) bool {
	switch compare.mode {
	case ByYear:
		return book1.Year() >= book2.Year()
	case BySize:
		return book1.Size() >= book2.Size()
	case ByRate:
		return book1.Rate() >= book2.Rate()
	default:
		return false
	}
}
