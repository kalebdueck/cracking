package chapter1

import "testing"

func TestUrlify(t *testing.T) {
	type args struct {
		s []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no spaces in here",
			args: args{s: []rune("yosup")},
			want: "yosup",
		},
		{
			name: "I found a space",
			args: args{s: []rune("yo sup")},
			want: "yo%20sup",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Urlify(tt.args.s); got != tt.want {
				t.Errorf("Urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
