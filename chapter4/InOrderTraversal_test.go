package chapter4

import (
	"reflect"
	"testing"
)

func TestInOrderTraversal(t *testing.T) {
	type args struct {
		root   *TreeNode
		sorted []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "balanced",
			args: args{
				root: &TreeNode{
					value: 5,
					left: &TreeNode{
						value: 3,
						left: &TreeNode{
							value: 1,
						},
						right: &TreeNode{
							value: 4,
						},
					},
					right: &TreeNode{
						value: 7,
					},
				},
			},
			want: []int{1, 3, 4, 5, 7},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrderTraversal(tt.args.root, tt.args.sorted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
