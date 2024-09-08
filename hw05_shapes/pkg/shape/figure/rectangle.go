package figure

import "fmt"

type Rectangle struct {
	width, height float64
}

func NewRectangle(width float64, height float64) *Rectangle {
	return &Rectangle{width, height}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Info() string {
	return fmt.Sprintf("Прямоугольник: ширина %.2f, высота %.2f", r.width, r.height)
}
