package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "a,a,a,a,a"

	if want != got {
		t.Errorf("got %q but want %q", got, want)
	}
}
