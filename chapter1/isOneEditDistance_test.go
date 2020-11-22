package chapter1

import "testing"

func TestIsOneEditDistance(t *testing.T) {
	type args struct {
		check   string
		compare string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "equal",
			args: args{
				check:   "pale",
				compare: "pale",
			},
			want: true,
		},
		{
			name: "down",
			args: args{
				check:   "ple",
				compare: "pale",
			},
			want: true,
		},
		{
			name: "up",
			args: args{
				check:   "pales",
				compare: "pale",
			},
			want: true,
		},
		{
			name: "swap",
			args: args{
				check:   "bale",
				compare: "pale",
			},
			want: true,
		},
		{
			name: "2swap",
			args: args{
				check:   "pale",
				compare: "bake",
			},
			want: false,
		},
		{
			name: "2down",
			args: args{
				check:   "pale",
				compare: "pa",
			},
			want: false,
		},
		{
			name: "2up",
			args: args{
				check:   "pale",
				compare: "palers",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneEditDistance(tt.args.check, tt.args.compare); got != tt.want {
				t.Errorf("IsOneEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
