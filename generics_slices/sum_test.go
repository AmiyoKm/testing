package slices

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d but want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{5, 7}, []int{0, 9})
		want := []int{7, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
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

	t.Run("sum of all elements", func(t *testing.T) {
		sum := func(x, y int) int {
			return x + y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, sum, 0), 6)
	})
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

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("Find the best programmer", func(t *testing.T) {
		type Person struct {
			Name string
		}
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

func TestMap(t *testing.T){
	t.Run("map through the slice and add 2",func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		addTwo := func(x int) int {
			return x + 2
		}
		got := Map(numbers,addTwo)
		if !slices.Equal(got, []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}) {
			t.Errorf("got %v, want %v", got, []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
		}
	})

	t.Run("map through the slice of string and add something",func(t *testing.T) {
		numbers := []string{"Anna", "Amiyo", "Ankita"}

		addHello := func(x string) string {
			return "Hello, " + x
		}
		got := Map(numbers,addHello)
		want :=  []string{"Hello, Anna", "Hello, Amiyo", "Hello, Ankita"}
		if !slices.Equal(got,want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("map through the slice of int and convert to string", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		toString := func(x int) string {
			return fmt.Sprintf("Number: %d", x)
		}

		got := Map(numbers, toString)
		want := []string{"Number: 1", "Number: 2", "Number: 3"}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertTrue(t testing.TB, condition bool) {
	t.Helper()
	if !condition {
		t.Errorf("expected true but got false")
	}
}
