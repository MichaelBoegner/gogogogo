package iteration

import "testing"

func TestRepeat(t *testing.T) {
	want := Repeat("a")
	got := "a,a,a,a,a"

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
