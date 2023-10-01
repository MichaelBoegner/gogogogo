package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("counter counts to three.", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertAmount(t, counter.Value(), 3)
	})
	t.Run("counter that's safe to run concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertAmount(t, counter.Value(), wantedCount)
	})
}

func assertAmount(t *testing.T, got, want int) {
	t.Helper()
	if got != 3 {
		t.Errorf("got %d, but wanted %d", got, want)
	}
}
