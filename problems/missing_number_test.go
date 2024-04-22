package problems

import "testing"

func TestMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 0",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "Case 1",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{0, 1, 3, 4, 5},
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{1, 2, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
