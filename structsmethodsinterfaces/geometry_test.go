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
	t.Run("Area() should take a height and width of a rectangle and return the area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})

}
