package problems

import (
	"reflect"
	"testing"
)

func TestCalculateDeliveryTime(t *testing.T) {
	type args struct {
		n int
		m int
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				n: 8,
				m: 0,
				t: 65,
			},
			want: "09:05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDeliveryTime(tt.args.n, tt.args.m, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateDeliveryTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
