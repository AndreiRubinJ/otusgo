package main

import (
	"fmt"

	"github.com/AndreiRubinJ/otusgo/hw05_shapes/pkg/shape"
	"github.com/AndreiRubinJ/otusgo/hw05_shapes/pkg/shape/figure"
)

func main() {
	cycle := figure.NewCircle(5)
	rectangle := figure.NewRectangle(5, 10)
	triangle := figure.NewTriangle(5, 10)
	emptyFigure := "dd"
	trapezoid := figure.NewTrapezoid(5, 5, 5)
	shapes := []any{cycle, rectangle, triangle, emptyFigure, trapezoid}
	for _, object := range shapes {
		if area, err := shape.CalculateArea(object); err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			printShapeInfo(area, object)
		}
	}
}

func printShapeInfo(area float64, s any) {
	shapeDate, ok := s.(shape.Info)
	if !ok {
		fmt.Println("Ошибка: объект не предоставляет информацию о форме")
		return
	}
	fmt.Printf("%s Площадь: %.2f\n", shapeDate.Info(), area)
}
