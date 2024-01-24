package recursive

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {
	t.Run("test gcd", func(t *testing.T) {
		got := GreatestCommonDivisor(10, 15)
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
