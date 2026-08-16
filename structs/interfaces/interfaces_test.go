package interfaces

import "testing"

func TestArea(t *testing.T){

	checkArea := func (t testing.TB, shape Shape, want float64)  {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g but want %g ", got, want)
		}
	}

	t.Run("Area of the Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12,6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("Area of Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 31.41592653589793)
	})

	t.Run("Area of Triangle", func(t *testing.T) {
		triangle := Triangle{12, 6}
		checkArea(t, triangle, 36.0)
	})
}

// Table driven tests

func TestAre(t *testing.T){

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 12, Length: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
        {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 31.41592653589793},
		// {Triangle{12, 6}, 31.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf(" %#v got %g want %g ", got, tt.want)
		}
	}
}

func TestAdvancedArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Length: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}