package problems

import (
	"reflect"
	"testing"
)

func TestFindQuadruples(t *testing.T) {
	type args struct {
		a1 int
		a2 int
		a3 int
		a4 int
		m  int
	}
	tests := []struct {
		name string
		args args
		want [][4]int
	}{
		{
			name: "Example 1",
			args: args{
				a1: 2,
				a2: 3,
				a3: 4,
				a4: 5,
				m:  30,
			},
			want: [][4]int{
				{1, 1, 5, 1},
				{1, 2, 3, 2},
				{1, 3, 1, 3},
				{1, 5, 2, 1},
				{2, 1, 2, 3},
				{2, 3, 3, 1},
				{2, 4, 1, 2},
				{3, 1, 4, 1},
				{3, 2, 2, 2},
				{3, 5, 1, 1},
				{4, 1, 1, 3},
				{4, 3, 2, 1},
				{5, 1, 3, 1},
				{5, 2, 1, 2},
				{6, 3, 1, 1},
				{7, 1, 2, 1},
				{9, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindQuadruples(tt.args.a1, tt.args.a2, tt.args.a3, tt.args.a4, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindQuadruples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindQuadruplesRecursive(t *testing.T) {
	type args struct {
		a               [4]int
		m               int
		currentSolution *[4]int
		index           int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case 1",
			args: args{
				a:               [4]int{2, 3, 4, 5},
				m:               30,
				currentSolution: nil,
				index:           0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindQuadruplesRecursive(tt.args.a, tt.args.m, tt.args.currentSolution, tt.args.index)
		})
	}
}
