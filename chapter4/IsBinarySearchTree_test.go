package chapter4

import "testing"

func TestIsBinarySearchTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "balanced cause empty",
			args: args{
				root: &TreeNode{
					value: 5,
				},
			},
			want: true,
		},
		{
			name: "unbalanced",
			args: args{
				root: &TreeNode{
					value: 5,
					left: &TreeNode{
						value: 3,
						left: &TreeNode{
							value: 1,
						},
						right: &TreeNode{
							value: 2,
						},
					},
					right: &TreeNode{
						value: 7,
					},
				},
			},
			want: false,
		},
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
			want: true,
		},
		{
			name: "unbalanced because a leaf is higher than a node more than 1 level up",
			args: args{
				root: &TreeNode{
					value: 5,
					left: &TreeNode{
						value: 3,
						left: &TreeNode{
							value: 1,
						},
						right: &TreeNode{
							value: 7,
						},
					},
					right: &TreeNode{
						value: 7,
					},
				},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBinarySearchTree(tt.args.root); got != tt.want {
				t.Errorf("IsBinarySearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
