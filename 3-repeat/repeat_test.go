package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("entering a should return aaaaaa", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("entering b should return bbbbbb", func(t *testing.T) {
		repeated := Repeat("b", 6)
		expected := "bbbbbb"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("entering 10 should repeat the count to 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 10)
	fmt.Println(repeated)
	// output: xxxxxxxxxx
}
