package hello

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello() = \"%s\", want \"%s\"", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Say 'Hello, World' is an empty string is supplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("")
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		args := "Kay"
		want := "Hello, Kay"
		got := Hello(args)
		if got != want {
			t.Errorf("Hello(%s) = \"%s\", want \"%s\"", args, got, want)
		}
	})
}
