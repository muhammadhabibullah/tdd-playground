package structs

import (
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type RightTriangle struct {
	Base   float64
	Height float64
}

func (rt RightTriangle) Area() float64 {
	return rt.Base * rt.Height / 2
}

func (rt RightTriangle) Perimeter() float64 {
	return rt.Base + rt.Height + math.Sqrt(rt.Base*rt.Base+rt.Height*rt.Height)
}
