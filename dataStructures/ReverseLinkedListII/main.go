package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// Function to print the linked list.
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func reverseBetween(head *Node, left int, right int) *Node {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	dummy := &Node{Next: head}
	prev := dummy

	// Move prev to the node before the left position
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Reverse the nodes between left and right
	current := prev.Next
	for i := left; i < right; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

func main() {
	// create a linked list
	head := &Node{Value: 1, Next: nil}
	node1 := &Node{Value: 2, Next: nil}
	node2 := &Node{Value: 3, Next: nil}
	node3 := &Node{Value: 4, Next: nil}
	node4 := &Node{Value: 5, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	left := 2
	right := 4
	// check if the linked list is a palindrome
	reversed := reverseBetween(head, left, right)
	printList(reversed)

}
