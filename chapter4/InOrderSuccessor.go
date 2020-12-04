package chapter4

func InOrderSuccessor(node *TreeNodeParentLink) *TreeNodeParentLink {

	//first node we look at could just have a right,
	//then just return the leftmost child

	if node.right != nil {
		return leftMostChild(node.right)
	}

	result := checkNextSuccessor(node)
	if result == nil {
		return node
	}

	return result
}
func leftMostChild(node *TreeNodeParentLink) *TreeNodeParentLink {
	for node.left != nil {
		node = node.left
	}

	return node
}

func checkNextSuccessor(node *TreeNodeParentLink) *TreeNodeParentLink {

	if node.parent == nil {
		return nil
	}

	if node.parent.value > node.value {
		return node.parent
	}

	return checkNextSuccessor(node.parent)

}
