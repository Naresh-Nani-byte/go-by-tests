package interfaces

import "math"


type Shape interface{
	Area() float64
}

type Rectangle struct{
	Length, Width float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 {
	return  r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height ) * 0.5
}