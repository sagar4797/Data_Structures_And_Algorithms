package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (head *Node) isPalindrome() bool {
	if head == nil {
		return true
	}

	// find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fmt.Println("fast:", fast.Value)
		fmt.Println("slow:", slow.Value)
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println("After fast:", fast.Value)
		fmt.Println("After slow:", slow.Value)
	}

	// reverse the second half of the linked list
	var prev *Node
	current := slow
	fmt.Println("Current", current.Value)
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	fmt.Println("Prev", prev)
	// compare the first half with the reversed second half
	for prev != nil && head != nil {
		if prev.Value != head.Value {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}

func main() {
	// create a linked list
	head := &Node{Value: 1, Next: nil}
	node1 := &Node{Value: 2, Next: nil}
	node2 := &Node{Value: 3, Next: nil}
	node3 := &Node{Value: 2, Next: nil}
	node4 := &Node{Value: 1, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// check if the linked list is a palindrome
	result := head.isPalindrome()
	fmt.Println("result", result) // should print "true"
}
