package chapter2

//kinda a silly one, we're only given access to the middle node. so we don't know head
func deleteMiddleNode(ll *nodeSingly) *nodeSingly {

	ll.val = ll.next.val
	ll.next = ll.next.next

	return ll
}
