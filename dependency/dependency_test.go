package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Inject dependency to read std.out for printed greeting from Greet()", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Mike")

		got := buffer.String()
		want := "Hello, Mike"

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})
}
