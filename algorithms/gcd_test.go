package algorithms

import "testing"

func TestNaiveGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gcd",
			args: args{
				a: 10,
				b: 15,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NaiveGCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("NaiveGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
