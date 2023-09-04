package slices

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	t.Run("SumSlice() take a slice and sum its contents", func(t *testing.T) {
		sliced := []int{1, 2, 3}
		got := SumSlice(sliced)
		want := 6

		if want != got {
			t.Errorf("wanted %d, but got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll() should take a varying number of slices, returning a new slice containing the totals for each slice passed in", func(t *testing.T) {

		got := SumAll([]int{1, 1}, []int{1, 2, 2})
		want := []int{2, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got back %v, wanted %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("SumAllTails should calculate the totals of the tails of each slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 1}, []int{1, 2, 2})
		want := 5

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but wanted %d", got, want)
		}
	})
}
