package figure

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) Info() string {
	return fmt.Sprintf("Круг: радиус %.2f", c.radius)
}
