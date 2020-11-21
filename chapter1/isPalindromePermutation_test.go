package chapter1

import "testing"

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is a palindrome",
			args: args{s: "racecar"},
			want: true,
		},
		{
			name: "isn't a palindrome",
			args: args{s: "racecars"},
			want: false,
		},
		{
			name: "is a palindrome permutation",
			args: args{s: "scaecarsr"},
			want: true,
		},
		{
			name: "is a palindrome permutation with spaces",
			args: args{s: " s cae   car   sr"},
			want: true,
		},
		{
			name: "empty string is I guess",
			args: args{s: ""},
			want: true,
		},
		{
			name: "1 char!",
			args: args{s: "!"},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
