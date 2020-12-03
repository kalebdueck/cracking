package chapter4

import "testing"

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "balanced",
			args: args{
				root: &TreeNode{
					value: 1,
					left:  nil,
					right: &TreeNode{
						value: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "unbalanced",
			args: args{
				root: &TreeNode{
					value: 1,
					left:  nil,
					right: &TreeNode{
						value: 2,
						left: &TreeNode{
							value: 3,
						},
					},
				},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.args.root); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
