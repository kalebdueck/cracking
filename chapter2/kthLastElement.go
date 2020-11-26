package chapter2

import (
	"fmt"
)

func kthLastElement(ll *nodeSingly, k int) *nodeSingly {
	a := ll
	b := ll
	var count int

	for a.next != nil {
		a = a.next
		count++
	} //We're gonna over set count here.

	fmt.Println(count)
	for i := 0; i <= count-k; i++ {
		b = b.next
	}

	return b
}
