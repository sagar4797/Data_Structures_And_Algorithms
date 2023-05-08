package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func numComponents(head *Node, nums []int) int {
	current := head
	isContinue := false
	count := 0

	for current != nil {
		for i := 0; i < len(nums); i++ {
			if current.Value == nums[i] {
				fmt.Println("current.Value ", current.Value)
				fmt.Println("nums[i] ", nums[i])
				if isContinue == false {
					count++
					isContinue = true
					continue
				}
			} else {
				isContinue = false
			}
		}

		current = current.Next
	}

	return count - 1
}

func main() {
	head := &Node{Value: 0, Next: nil}
	node1 := &Node{Value: 1, Next: nil}
	node2 := &Node{Value: 2, Next: nil}
	node3 := &Node{Value: 3, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	nums := []int{0, 1, 3}
	count := numComponents(head, nums)
	fmt.Println("Count ", count)
}
