package chapter4

//InOrderTraversal prints a binary tree in order
func InOrderTraversal(root *TreeNode, sorted []int) []int {

	if root == nil {
		return sorted
	}

	sorted = InOrderTraversal(root.left, sorted)
	sorted = append(sorted, root.value)
	sorted = InOrderTraversal(root.right, sorted)

	return sorted
}
