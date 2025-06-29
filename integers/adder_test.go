package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expected := 4

	assertCorrectSum(t, got, expected)
}

func assertCorrectSum(t testing.TB, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %d but go %d", expected, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
