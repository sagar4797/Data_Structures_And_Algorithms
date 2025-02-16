package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree() {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := 0
	maxID := -1

	for idx, num := range nums {
		if maxID < 0 || num > max {
			max = num
			maxID = idx
		}
	}

	newNode := TreeNode{max, nil, nil}

	if maxID > 0 {
		newNode.Left = constructMaximumBinaryTree(nums[:maxID])
	}

	if maxID < len(nums)-1 {
		newNode.Right = constructMaximumBinaryTree(nums[maxID+1:])
	}
	return &newNode
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	treeNode := constructMaximumBinaryTree(nums)
	fmt.Println(treeNode)
}
