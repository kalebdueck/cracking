package chapter3

import (
	"reflect"
	"testing"
)

func TestSortStack(t *testing.T) {
	type args struct {
		stack []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "full sort",
			args: args{
				stack: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "jumble sort",
			args: args{
				stack: []int{5, 2, 3, 4, 1},
			},
			want: []int{5, 4, 3, 2, 1},
		},

		{
			name: "no sort",
			args: args{
				stack: []int{5, 4, 3, 2, 1},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortStack(tt.args.stack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
