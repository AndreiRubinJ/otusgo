package shape

import "errors"

type Shape interface {
	Area() float64
}

type Info interface {
	Info() string
}

func CalculateArea(s interface{}) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}
