package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a",10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeated(b *testing.B) {
	for b.Loop() {
		Repeat("a",20)
	}
}

func ExampleRepeat() {
	res := Repeat("b", 5)
	fmt.Println(res)
	// Output: bbbbb
}