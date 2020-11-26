package chapter2

import (
	"fmt"
)

func removeDuplicates(ll *nodeSingly) *nodeSingly {

	seenVals := make(map[string]bool)
	n := ll

	for n.next != nil {
		fmt.Printf("start: %s\n", n.val)

		seenVals[n.val] = true

		if seenVals[n.next.val] == true {
			n.next = n.next.next
		} else {
			n = n.next
		}
	}

	return ll
}

func removeDuplicatesNoSpace(ll *nodeSingly) *nodeSingly {
	current := ll

	for current != nil && current.next != nil {
		cast := current
		for cast.next != nil {
			if current.val == cast.next.val {
				cast.next = cast.next.next
			} else {
				cast = cast.next
			}
		}

		current = current.next
	}

	return ll
}
