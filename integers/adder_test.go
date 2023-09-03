package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(8, 2)
	expected := 10

	if sum != expected {
		t.Errorf("expected %d, but got %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(10, 2)
	fmt.Println(sum)
	// Output: 12
}
