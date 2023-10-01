package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("counter that is safe to use concurrently.", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		if counter.Value() != 3 {
			t.Errorf("got %d, but wanted %d", counter.Value(), 3)
		}

	})
}
