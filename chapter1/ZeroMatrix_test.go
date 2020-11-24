package chapter1

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3x3",
			args: args{
				matrix: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
		},
		{
			name: "zero in two rows",
			args: args{
				matrix: [][]int{
					[]int{0, 1, 2, 3},
					[]int{1, 5, 6, 1},
					[]int{0, 8, 9, 1},
				},
			},
			want: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 5, 6, 1},
				[]int{0, 0, 0, 0},
			},
		},
		{
			name: "zero everywhere",
			args: args{
				matrix: [][]int{
					[]int{0, 1, 2},
					[]int{0, 5, 6},
					[]int{0, 8, 9},
				},
			},
			want: [][]int{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
		{
			name: "one col",
			args: args{
				matrix: [][]int{
					[]int{0, 1, 2},
					[]int{1, 5, 6},
					[]int{2, 8, 9},
				},
			},
			want: [][]int{
				[]int{0, 0, 0},
				[]int{0, 5, 6},
				[]int{0, 8, 9},
			},
		},
		{
			name: "four corners",
			args: args{
				matrix: [][]int{
					[]int{0, 1, 2, 1, 0},
					[]int{1, 5, 6, 1, 1},
					[]int{0, 8, 9, 1, 0},
				},
			},
			want: [][]int{
				[]int{0, 0, 0, 0, 0},
				[]int{0, 5, 6, 1, 0},
				[]int{0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
