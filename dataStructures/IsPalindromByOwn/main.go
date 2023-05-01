package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 1, Next: nil}
	node1 := &Node{Value: 2, Next: nil}
	node2 := &Node{Value: 3, Next: nil}
	node3 := &Node{Value: 4, Next: nil}
	node4 := &Node{Value: 1, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	result := head.IsPalindrom()
	fmt.Println("Result:", result)

}

func (head *Node) IsPalindrom() bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fmt.Println("fast:", fast.Value)
		fmt.Println("slow:", slow.Value)
		slow = slow.Next
		fast = fast.Next.Next
	}

	current := slow
	var prev *Node

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	for prev != nil && head != nil {
		if prev.Value != head.Value {
			return false
		}

		prev = prev.Next
		head = head.Next
	}
	return true
}

// 1 ->2 -> 3 -> 2 -> 1
