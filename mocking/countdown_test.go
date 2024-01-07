package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("buffer.String() = \"%v\"; want \"%v\"", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("spySleeper wasn't called %d times", 3)
	}
}
