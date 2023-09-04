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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{shape: Circle{Radius: 5.0}, want: 78.53981633974999},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, but wanted %g", got, tt.want)
		}
	}
}
