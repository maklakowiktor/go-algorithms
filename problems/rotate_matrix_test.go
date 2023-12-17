package problems

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Common matrix",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
