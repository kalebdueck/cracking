package chapter1

import "testing"

func Test_isPermutation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is a permutation",
			args: args{
				a: "hello",
				b: "elloh",
			},
			want: true,
		},
		{
			name: "is not a permutation",
			args: args{
				a: "hello",
				b: "ellhh",
			},
			want: false,
		},

		{
			name: "is empty",
			args: args{
				a: "",
				b: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPermutation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
