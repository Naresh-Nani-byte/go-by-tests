package structs

import (
	"testing"
)

func TestShapesPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := ShapesPerimeter(rectangle)
	want := 40.0

	if got != want{
		t.Errorf("got %.2f want %.2f ", got, want)
	}
}

func TestShapesArea(t *testing.T){
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 2.0}
		got := ShapesArea(rectangle)
		want := 10.0

		if got != want{
			t.Errorf("got %g want %g ", got, want)
		}
	})

	// t.Run("Circle ", func(t *testing.T) {
	// 	circle := Circle{1.7}
	// 	got := ShapesArea(circle)
	// 	want := math.Pi * 10
	// 	if got != want{
	// 		t.Errorf("got %g want %g ", got, want)
	// 	}
	// })
	
}
