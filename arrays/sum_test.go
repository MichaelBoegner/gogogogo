package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum will take an array of numbers and return the total", func(t *testing.T) {
		array := [3]int{1, 2, 3}
		got := Sum(array)
		want := 6

		if got != want {
			t.Errorf("Got %d back, but wanted %d", got, want)
		}
	})

}
