package chapter4

func IsBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return IsBinarySearchTreeMinMax(root.left, 0, root.value) && IsBinarySearchTreeMinMax(root.right, root.value, 0)
}

func IsBinarySearchTreeMinMax(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if (min != 0 && root.value <= min) || (max != 0 && root.value >= max) {
		return false
	}

	return IsBinarySearchTreeMinMax(root.left, min, root.value) && IsBinarySearchTreeMinMax(root.right, root.value, max)

}
