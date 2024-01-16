package misc

import (
	"reflect"
	"testing"
)

func TestSubtractSlice(t *testing.T) {
	t.Run("subtract int slice", func(t *testing.T) {
		got := SubtractSlice([]int{1, 2, 3, 4}, []int{2, 3})
		want := []int{1, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("subtract string slice", func(t *testing.T) {
		got := SubtractSlice([]string{"1", "2", "3", "4"}, []string{"2", "3"})
		want := []string{"1", "4"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("subtract empty slice", func(t *testing.T) {
		got := SubtractSlice([]string{"2", "3"}, []string{})
		want := []string{"2", "3"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("subtract slice from empty slice", func(t *testing.T) {
		got := SubtractSlice([]string{}, []string{"2", "3"})
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

func TestIntersectSlices(t *testing.T) {
	t.Run("intersect int slices", func(t *testing.T) {
		got := IntersectSlices([]int{1, 2, 3, 4}, []int{2, 3})
		want := []int{2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("intersect string slices", func(t *testing.T) {
		got := IntersectSlices([]string{"1", "2", "3", "4"}, []string{"2", "3"})
		want := []string{"2", "3"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("intersect empty slice", func(t *testing.T) {
		got := IntersectSlices([]string{"2", "3"}, []string{})
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("intersect slice from empty slice", func(t *testing.T) {
		got := IntersectSlices([]string{}, []string{"2", "3"})
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

func TestUniqueSlices(t *testing.T) {
	t.Run("two same int elements", func(t *testing.T) {
		got := UniqueSlice([]int{1, 2, 2, 4})
		want := []int{1, 2, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("multiple same string elements", func(t *testing.T) {
		got := UniqueSlice([]string{"1", "1", "1", "1", "4"})
		want := []string{"1", "4"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("no same string elements", func(t *testing.T) {
		got := UniqueSlice([]string{"a1b2c3", "0000a2222cccc3333"})
		want := []string{"a1b2c3", "0000a2222cccc3333"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("empty slices", func(t *testing.T) {
		got := UniqueSlice([]string{})
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}
