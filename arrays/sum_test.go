package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15
		if got != want {
			t.Errorf("Sum(%v) = \"%d\"; want \"%d\"", arr, got, want)
		}
	})
}
