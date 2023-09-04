package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter() should take a width and height of a rectangle and calculate the perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, but wanted %g", got, want)
		}
	}

	t.Run("Area(rectangle Rectangle) should take a height and width of a rectangle and return the area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("Area(circle Circle) should take the radius of a circle and return the area", func(t *testing.T) {
		circle := Circle{5.0}
		checkArea(t, circle, 78.53981633974999)
	})
}
