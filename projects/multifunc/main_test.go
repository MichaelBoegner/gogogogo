package multifunc

import "testing"

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSubtract(t *testing.T) {

	got := Subtract(10, 20)
	want := -1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
