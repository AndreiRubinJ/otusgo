package figure

import "fmt"

type Triangle struct {
	base, height float64
}

func NewTriangle(base float64, height float64) *Triangle {
	return &Triangle{base, height}
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t *Triangle) Info() string {
	return fmt.Sprintf("Треугольник: основание %.2f, высота %.2f", t.base, t.height)
}
