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

	got := SumAll(nbrSlices...)
	want := []int{15, 18}
	assertEqualSlices(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Expected Case", func(t *testing.T) {
		nbrSlices := [][]int{
			{1, 2, 3, 4, 5},
			{5, 6, 7},
		}

		got := SumAllTails(nbrSlices...)
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
			got := SumAllTails(tc.input...)
			assertEqualSlices(t, got, tc.want)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func AssertEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		t.Errorf("a = %v not equal b = %v", a, b)
	}
}

func AssertNotEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf("a = %v equal b = %v", a, b)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got is false")
	}
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func BalanceFor(txs []Transaction, name string) float64 {
	sumTxs := func(currBal float64, t2 Transaction) float64 {
		if t2.From == name {
			return currBal - t2.Sum
		} else if t2.To == name {
			return currBal + t2.Sum
		}
		return currBal
	}

	return Reduce(txs, sumTxs, 0)
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}
