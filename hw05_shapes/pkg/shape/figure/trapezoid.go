package figure

type Trapezoid struct {
	width, height, radius float64
}

func NewTrapezoid(radius float64, height float64, width float64) *Trapezoid {
	return &Trapezoid{width, height, radius}
}
