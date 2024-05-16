package main

import (
	"fmt"
)

type Shape interface {
	Name() string
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

func main() {
	var s Shape
	s = Rectangle{
		Length: 10,
		Width:  5,
	}
	ShapeInformation(s)

	s = Circle{
		Radius: 12,
	}
	ShapeInformation(s)

	s = Square{
		Length: 9}
	ShapeInformation(s)
}

func ShapeInformation(s Shape) {
	fmt.Printf("Area of %s: %.2f\n", s.Name(), s.Area())
	fmt.Printf("Perimeter of %s: %.2f\n", s.Name(), s.Perimeter())

}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.Length) + (2 * r.Width)
}

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.141592654
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.141592654

}

func (s Square) Name() string {
	return "Square"
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}
