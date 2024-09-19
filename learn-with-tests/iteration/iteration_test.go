package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times letter a", func(t *testing.T) {
		result := Repeat("a", 5)
		expected := "aaaaa"

		if (result != expected) {
			t.Errorf("got %q want %q", result, expected)
		}
	})
	t.Run("repeat 5 times letter b", func(t *testing.T) {
		result := Repeat("b", 5)
		expected := "bbbbb"

		if (result != expected) {
			t.Errorf("got %q want %q", result, expected)
		}
	})
	t.Run("repeat 5 times letter c", func(t *testing.T) {
		result := Repeat("c", 5)
		expected := "ccccc"

		if (result != expected) {
			t.Errorf("got %q want %q", result, expected)
		}
	})

	t.Run("repeat 4 times letter c", func(t *testing.T) {
		result := Repeat("c", 4)
		expected := "cccc"

		if (result != expected) {
			t.Errorf("got %q want %q", result, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}