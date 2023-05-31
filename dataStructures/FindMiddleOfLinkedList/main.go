package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (head *Node) middleNode() *Node {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	head := Node{Value: 1}
	head1 := Node{Value: 2}
	head2 := Node{Value: 3}
	head3 := Node{Value: 4}
	head4 := Node{Value: 5}
	head5 := Node{Value: 6}

	head.Next = &head1
	head1.Next = &head2
	head2.Next = &head3
	head3.Next = &head4
	head4.Next = &head5

	middleNode := head.middleNode()
	fmt.Printf("This is middle node of the linked list: %+v", middleNode)
}
