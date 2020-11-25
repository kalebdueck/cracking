package chapter2

func removeDuplicates(ll *nodeSingly) *nodeSingly {

	seenVals := make(map[string]bool)
	seenVals[ll.val] = true
	n := ll

	for n.next != nil {

		if seenVals[n.next.val] == true {
			n.next = n.next.next
		}

		seenVals[n.val] = true

		n = n.next
	}

	return ll
}

func removeDuplicatesNoSpace(ll *nodeSingly) *nodeSingly {

	seenVals := make(map[string]bool)
	seenVals[ll.val] = true
	current := ll
	cast := ll.next

	for current.next != nil {
		for cast.next != nil {
			if current.val == cast.val {
				current.next = current.next.next
			}

			cast = cast.next
		}

		current = current.next
		cast = current.next
	}

	return ll
}
