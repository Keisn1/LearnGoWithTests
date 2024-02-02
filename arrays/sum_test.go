package arrays

import (
	"reflect"
	"strings"
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
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func NewBalanceFor(account Account, txs []Transaction) Account {
	sumTxs := func(account Account, tx Transaction) Account {
		if tx.From == account.Name {
			account.Balance -= tx.Amount
		} else if tx.To == account.Name {
			account.Balance += tx.Amount
		}
		return account
	}

	return Reduce(txs, sumTxs, account)
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{
		From:   from.Name,
		To:     to.Name,
		Amount: amount,
	}
}

func BalanceFor(txs []Transaction, name string) float64 {
	sumTxs := func(currBal float64, t2 Transaction) float64 {
		if t2.From == name {
			return currBal - t2.Amount
		} else if t2.To == name {
			return currBal + t2.Amount
		}
		return currBal
	}

	return Reduce(txs, sumTxs, 0)
}

type Transaction struct {
	From   string
	To     string
	Amount float64
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})

}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
