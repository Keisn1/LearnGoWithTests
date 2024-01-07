package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("Test happy path and amount calling", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("buffer.String() = \"%v\"; want \"%v\"", got, want)
		}
	})

	t.Run("Test order of Countdown operations", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}
		Countdown(spyCountdownOperations, spyCountdownOperations)
		wantCalls := []string{
			write, sleep,
			write, sleep,
			write, sleep,
			write}

		if !slices.Equal(spyCountdownOperations.Calls, wantCalls) {
			t.Errorf("Not the right order")
			t.Error(spyCountdownOperations.Calls)
		}
	})

	t.Run("Test configurable Sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
