package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(8, 2)
	expected := 10

	if sum != expected {
		t.Errorf("expected %d, but got %d", sum, expected)
	}
}
