package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeatd := Repeat("a")
	expected := "aaaaa"

	if repeatd != expected {
		t.Errorf("expected %q but got %q", expected, repeatd)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
