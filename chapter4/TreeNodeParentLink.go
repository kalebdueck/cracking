package chapter4

type TreeNodeParentLink struct {
	value  int
	left   *TreeNodeParentLink
	right  *TreeNodeParentLink
	parent *TreeNodeParentLink
}
