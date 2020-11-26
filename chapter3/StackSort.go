package chapter3

import (
	"fmt"
)

func SortStack(stack []int) []int {
	var temp []int

	for isEmpty(stack) == false {
		var val int
		var curr int
		if isEmpty(temp) {
			val, stack = pop(stack)
			temp = push(val, temp)
		}

		curr, stack = pop(stack)
		fmt.Println(curr)
		fmt.Println(stack)

		for peek(temp) > curr {
			fmt.Println("in")
			val, temp = pop(temp)
			stack = push(val, stack)
		}

		fmt.Println("out")

		temp = push(curr, temp)
	}

	var val int
	for isEmpty(temp) == false {
		val, temp = pop(temp)
		stack = push(val, stack)
	}

	return stack
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func push(val int, stack []int) []int {
	return append(stack, val)
}

func peek(stack []int) int {
	if len(stack) == 0 {
		return 0
	}

	return stack[len(stack)-1]
}

func isEmpty(stack []int) bool {
	if len(stack) == 0 {
		return true
	}

	return false
}
