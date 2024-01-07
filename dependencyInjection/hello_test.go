package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("With buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("buffer.String() = \"%v\"; want \"%v\"", got, want)
		}
	})

	t.Run("With stdout", func(t *testing.T) {
		got, err := Greet(os.Stdout, "Chris")

		if err != nil {
			t.Errorf("Didn't want error, got %q", err)
		}

		want := len("Hello, Chris")

		if got != want {
			t.Errorf("buffer.String() = \"%v\"; want \"%v\"", got, want)
		}
	})
}
