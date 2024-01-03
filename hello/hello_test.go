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
		lang := "en"
		want := "Hello, World"
		got := Hello("", lang)
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Kay"
		lang := "en"
		want := "Hello, Kay"
		got := Hello(name, lang)
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Kay"
		lang := "asdf"
		want := "Hello, Kay"
		got := Hello(name, lang)
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in spanish", func(t *testing.T) {
		name := "Kay"
		lang := "esp"
		want := "Buenas Dias, Kay"
		got := Hello(name, lang)
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in french", func(t *testing.T) {
		name := "Kay"
		lang := "fr"
		want := "Bonjour, Kay"
		got := Hello(name, lang)
		assertCorrectMessage(t, got, want)
	})
}
