package methods

import "testing"

func TestShapes(t *testing.T){
	t.Run("Area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("Area of Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 31.41592653589793
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}