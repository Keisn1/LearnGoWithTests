package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Repeat() = \"%s\"; want \"%s\"", got, want)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("Printing n times the same character", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrectString(t, got, want)
	})
}

func ExampleRepeat() {
	got := Repeat("Again")
	fmt.Println(got)

	// Output:
	// AgainAgainAgainAgainAgain
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("Again")
	}
}
