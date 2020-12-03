package chapter4

//IsBalanced using a breadth first search. Seems workable, needs testing
//Book recommends a recursive bottom-up solution
//that needs to be replicated
func IsBalanced(root *TreeNode) bool {
	if checkHeight(root) == -1 {
		return false
	}

	return true
}

func checkHeight(root *TreeNode) int {
	if root == nil {
		return 0 //height of 0
	}

	leftHeight := checkHeight(root.left)
	rightHeight := checkHeight(root.right)

	if leftHeight == -1 || rightHeight == -1 {
		return -1 //Bubble up unbalance
	}

	heightDiff := leftHeight - rightHeight
	if heightDiff < 0 {
		heightDiff *= -1
	}
	if heightDiff > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}
