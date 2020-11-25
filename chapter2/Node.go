package chapter2

type nodeSingly struct {
	next *nodeSingly
	val  string
}

type nodeDoubly struct {
	next *nodeSingly
	prev *nodeSingly
	val  string
}
