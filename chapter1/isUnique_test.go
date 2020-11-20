package chapter1

import "testing"

func Test_isUnique(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple unique string returns true",
			args: args{s: "hiya bud"},
			want: true,
		},
		{
			name: "simple non-unique string returns false",
			args: args{s: "hiya buddy"},
			want: false,
		},
		{
			name: "zero length string is unique i guess",
			args: args{s: ""},
			want: true,
		},
		{
			name: "one char length string is unique",
			args: args{s: "a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.s); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}

			if got := isUniqueSavingSpace(tt.args.s); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
