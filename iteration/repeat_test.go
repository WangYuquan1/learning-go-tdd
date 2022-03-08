package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeatd := Repeat("a", 5)
	expected := "aaaaa"

	if repeatd != expected {
		t.Errorf("expected %q but got %q", expected, repeatd)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Printf("%q", repeated)
	// Output: aaaaa
}
