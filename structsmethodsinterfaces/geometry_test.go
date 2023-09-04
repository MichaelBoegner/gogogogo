package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter() should take a width and height of a rectangle and calculate the perimeter", func(t *testing.T) {
		got := Perimeter(15, 10)
		want := 50

		if got != want {
			t.Errorf("got %d, but wanted %d", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Area() should take a height and width of a rectangle and return the area", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f, but wanted %.2f", got, want)
		}
	})

}
