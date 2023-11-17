package algorithms

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Run("empty collection", func(t *testing.T) {
		got := []int{}
		InsertionSort(&got, 0)
		want := []int{}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("one-element collection", func(t *testing.T) {
		got := []int{0}
		InsertionSort(&got, 1)
		want := []int{0}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("collection of 4 numbers", func(t *testing.T) {
		got := []int{1, 2, 3, 4}
		InsertionSort(&got, 4)
		want := []int{4, 3, 2, 1}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
