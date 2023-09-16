package buffer

import "testing"

func TestBuffering(t *testing.T) {
	got := Buffering()
	want := "i don't know"

	if got != want {
		t.Errorf("got %v, but wanted %v", got, want)
	}

}
