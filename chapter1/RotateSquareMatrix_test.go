package chapter1

import (
	"reflect"
	"testing"
)

func TestRotateSquareMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
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
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateSquareMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateSquareMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
