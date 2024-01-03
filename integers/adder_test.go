package integers

import (
	"fmt"
	"testing"
)

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Adder() = \"%d\", want \"%d\"", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("2 plus 2 = 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectResult(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Printf("The result of 2 + 2 is %d", sum)

	// Output:
	// The result of 2 + 2 is 4
}
