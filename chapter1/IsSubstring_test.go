package chapter1

import "testing"

func TestIsSubstring(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is prefixed",
			args: args{
				s1: "you",
				s2: "yo",
			},
			want: true,
		},
		{
			name: "is postfixed",
			args: args{
				s1: "you",
				s2: "ou",
			},
			want: true,
		},
		{
			name: "is in the middle",
			args: args{
				s1: "your",
				s2: "ou",
			},
			want: true,
		},
		{
			name: "is ==",
			args: args{
				s1: "your",
				s2: "your",
			},
			want: true,
		},
		{
			name: "too long",
			args: args{
				s1: "you",
				s2: "your",
			},
			want: false,
		},
		{
			name: "no match",
			args: args{
				s1: "your",
				s2: "yor",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubstring(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
