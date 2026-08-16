package structs

type Rectangle struct{
	Width float64
	Length float64
}


func ShapesPerimeter(rectangle Rectangle) float64 {
	return  2 * (rectangle.Length + rectangle.Width)
}


func ShapesArea(rectangle Rectangle) float64 {
	return  rectangle.Length * rectangle.Width
}