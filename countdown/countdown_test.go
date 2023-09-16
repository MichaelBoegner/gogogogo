package countdown

import (
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3", func(t *testing.T) {

		Countdown(buffer)

		got := buffer.String()
		want := "3"

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})
}
