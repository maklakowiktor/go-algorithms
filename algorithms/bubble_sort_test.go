package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("empty collection", func(t *testing.T) {
		got := []int{}
		BubbleSort(&got, 0)
		want := []int{}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("one-element collection", func(t *testing.T) {
		got := []int{0}
		BubbleSort(&got, 1)
		want := []int{0}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("collection of 4 numbers", func(t *testing.T) {
		got := []int{4, 3, 2, 1}
		BubbleSort(&got, 4)
		want := []int{1, 2, 3, 4}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
