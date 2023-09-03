package slices

import "testing"

func TestSlice(t *testing.T) {
	t.Run("take a slice and sum its contents", func(t *testing.T) {
		sliced := []int{1, 2, 3}
		got := Slice(sliced)
		want := 6

		if want != got {
			t.Errorf("wanted %d, but got %d", want, got)
		}
	})
}
