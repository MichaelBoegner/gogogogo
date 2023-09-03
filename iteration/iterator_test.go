package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("single letter, no determined number of repeats", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "a,a,a,a,a"

		if want != got {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("caller determines number of repeats", func(t *testing.T) {
		got := Repeat("a", 2)
		want := "a,a"

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	zeroCalled := Repeat("a", 0)
	fmt.Println(zeroCalled)

	twoCalled := Repeat("c", 2)
	fmt.Println(twoCalled)
	// Output:
	//a,a,a,a,a
	//c,c
}
