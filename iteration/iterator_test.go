package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("single letter", func(t *testing.T) {
		got := Repeat("a")
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
		Repeat("a")
	}
}
