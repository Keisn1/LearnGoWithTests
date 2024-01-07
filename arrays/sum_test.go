package arrays

import (
	"reflect"
	"testing"
)

func assertEqualSlices(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SumAll() = \"%v\"; want \"%v\"", got, want)
	}

}

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

func TestSumAll(t *testing.T) {
	nbrSlices := [][]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7},
	}

	got := SumAll(nbrSlices)
	want := []int{15, 18}
	assertEqualSlices(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Expected Case", func(t *testing.T) {
		nbrSlices := [][]int{
			{1, 2, 3, 4, 5},
			{5, 6, 7},
		}

		got := SumAllTails(nbrSlices)
		want := []int{14, 13}
		assertEqualSlices(t, got, want)
	})

	t.Run("Edge Cases", func(t *testing.T) {
		testCases := []struct {
			input [][]int
			want  []int
		}{
			{
				input: [][]int{{}, {5, 6, 7}},
				want:  []int{0, 13},
			},
			{
				input: [][]int{{}, {}},
				want:  []int{0, 0},
			},
			{
				input: [][]int{{4}, {}},
				want:  []int{0, 0},
			},
		}

		for _, tc := range testCases {
			got := SumAllTails(tc.input)
			assertEqualSlices(t, got, tc.want)
		}
	})
}
