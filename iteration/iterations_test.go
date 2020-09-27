package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeatString := Repeat("a", 5)
	fmt.Println(repeatString)
	// Output: aaaaa
}
