package geometry

import (
	"fmt"
	"math"
)

type square struct {
	x1,
	y1,
	x2,
	y2,
	x3,
	y3,
	x4,
	y4 float64
}

func NewSquare(x1, y1, x2, y2, x3, y3, x4, y4 float64) square {
	fmt.Println("New Square")
	return square{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
		x3: x3,
		y3: y3,
		x4: x4,
		y4: y4,
	}
}

func (s square) getArea() {
	area := math.Pow(Length(s.x1, s.y1, s.x2, s.y2), 2)
	fmt.Println("Square Area:", area)
}

func (s square) getPerimeter() {
	p := Length(s.x1, s.y1, s.x2, s.y2) * 4
	fmt.Println("Square Perimeter:", p)
}
