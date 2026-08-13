package structs

import "testing"

func TestRectangle(t *testing.T){
	t.Run("testing the Perimeter of the rectangle", func(t *testing.T) {
		got := Perimeter(4.0,8.0)
		want := 24.0
		assertMessage(t, got, want)
	})

	t.Run("Test the Area of the triangle", func(t *testing.T) {
		got := Area(4.0,8.0)
		want := 32.0
		assertMessage(t, got, want)
	})
}


func assertMessage(t testing.TB, got, want float64){
	t.Helper()
	if got != want{
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}