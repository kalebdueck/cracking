package chapter1

import "testing"

func TestIsStringRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple rotation",
			args: args{
				s1: "waterbottle",
				s2: "erbottlewat",
			},
			want: true,
		},
		{
			name: "no rotation",
			args: args{
				s1: "waterbottle",
				s2: "erboltlewat",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsStringRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
