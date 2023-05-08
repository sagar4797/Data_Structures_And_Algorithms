package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	Tree := &Node{Val: 5, Left: &Node{Val: 3, Left: &Node{Val: 2, Left: nil, Right: nil}, Right: &Node{Val: 4, Left: nil, Right: nil}}, Right: &Node{Val: 9, Left: &Node{Val: 7, Left: nil, Right: nil}, Right: &Node{Val: 10, Left: nil, Right: nil}}}
	Tree.InOrderTreeTravalsal()
	Tree.Insert(8)
	fmt.Println(" ")
	Tree.InOrderTreeTravalsal()
}

func (root *Node) Insert(val int) *Node {
	if root == nil {
		return &Node{Val: val, Left: nil, Right: nil}

	}

	if val < root.Val {
		node := root.Left.Insert(val)
		root.Left = node
	}

	if val > root.Val {
		node := root.Right.Insert(val)
		root.Right = node
	}

	return root
}

func (root *Node) InOrderTreeTravalsal() {
	if root == nil {
		return
	}

	root.Left.InOrderTreeTravalsal()
	fmt.Printf("%d", root.Val)
	root.Right.InOrderTreeTravalsal()
}
