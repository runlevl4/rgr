package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
