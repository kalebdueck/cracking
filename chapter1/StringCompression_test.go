package chapter1

import "testing"

func TestCompressString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "is all ones",
			args: args{s: "abcdefg"},
			want: "abcdefg",
		},
		{
			name: "buncha dupes",
			args: args{s: "aaaabbbbcccdefg"},
			want: "a4b4c3d1e1f1g1",
		},
		{
			name: "same char in different spots",
			args: args{s: "aaaabbbbcccdefgaabba"},
			want: "a4b4c3d1e1f1g1a2b2a1",
		},
		{
			name: "one char",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "zero char",
			args: args{s: ""},
			want: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString(tt.args.s); got != tt.want {
				t.Errorf("CompressString() = %v, want %v", got, tt.want)
			}

			if got := CompressStringBookMethod(tt.args.s); got != tt.want {
				t.Errorf("CompressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
