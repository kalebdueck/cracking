package chapter4

func ListOfDepths(root *TreeNode) [][]*TreeNode {
	result := make([][]*TreeNode, 0)

	current := []*TreeNode{root}

	for len(current) > 0 {
		//add previous level
		result = append(result, current)
		//go to next level
		parents := current

		current = []*TreeNode{}

		for _, parent := range parents {

			if parent.left != nil {
				current = append(current, parent.left)
			}

			if parent.right != nil {
				current = append(current, parent.right)
			}
		}

	}

	return result
}
