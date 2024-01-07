package shapes

import "math"

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Perimeter() float64 {
	return 0.5 * (t.Base * t.Height)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
