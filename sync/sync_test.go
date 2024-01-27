package main

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		assertCount(t, counter, want)
	})

	t.Run("Increment counter in concurrent Environment", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < 1000; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}
		wg.Wait()

		time.Sleep(100 * time.Millisecond)
		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("value is %d; want %d", got.Value(), want)
	}
}
