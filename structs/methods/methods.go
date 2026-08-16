package methods

import "math"

type Rectangle struct{
	Width float64
	Length float64
}

type Circle struct {
	Radius float64
}


func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return  math.Pi * c.Radius
}

