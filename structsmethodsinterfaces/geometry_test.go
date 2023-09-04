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
