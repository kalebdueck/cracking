package chapter4

func balanceSortedArray(sortedArray []int) *TreeNode {
  //Do i need to think about odd/even?
  middle := len(sortedArray) / 2

  root := TreeNode{
    value: sortedArray[middle],
    left: balanceSortedArray(sortedArray[:middle]),
    right: balanceSortedArray(sortedArray[middle+1:]),
  }

  return &root
}
