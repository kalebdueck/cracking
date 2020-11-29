package chapter4

//Todo needs a DirectedGraph struct?
//Then two node structs
//We'll use a Breadth-first search initially (Can use a bidirectional search if asked)
//Then just exit out if we find the node we're looking for.
func RouteBetweenNodes(root *Node, search *Node) bool {
	//We'll use an array as a queue 
	// cause I'm too lazy to write a Queue class again
	queue := []*Node

	queue = append(queue, root)

	for len(queue) > 0 {
		//dequeue root
		node, queue := queue[0], queue[1:]

		//If we've been there already, we can walk away
		if node.Visited {
			continue
		}

		//check if node is search
		if node.Name == search.Name {
			return true
		}

		node.Visited = true

		queue = append(queue, node.Children)

	}

	return false
}
