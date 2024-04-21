package slidingwindow

import (
	"reflect"
	"testing"
)

func TestAverageOfSubarrayOfSizeK(t *testing.T) {
	type args struct {
		items []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Average int slice with negative item and K = 5",
			args: args{
				items: []int{1, 3, 2, 6, -1, 4, 1, 8, 2},
				k:     5,
			},
			want: []float64{2.2, 2.8, 2.4, 3.6, 2.8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Fix got != tt.want
			if got := AverageOfSubarrayOfSizeK(tt.args.items, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AverageOfSubarrayOfSizeK() = %v, want %v", got, tt.want)
			}
		})
	}
}
