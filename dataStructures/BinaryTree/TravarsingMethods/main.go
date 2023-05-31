package main

import "fmt"

// Preoder -> root -> left-> right
// Post order -> left -> right -> root
//Inorder -> left -> root -> right

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	Tree := &Node{Val: 5, Left: &Node{Val: 3, Left: &Node{Val: 2, Left: nil, Right: nil}, Right: &Node{Val: 4, Left: nil, Right: nil}}, Right: &Node{Val: 9, Left: &Node{Val: 7, Left: nil, Right: nil}, Right: &Node{Val: 10, Left: nil, Right: nil}}}
	fmt.Println("Inside main")
	Tree.PreOrderTreeTravalsal()
	fmt.Println(" ")
	Tree.PostOrderTreeTravalsal()
	fmt.Println(" ")
	Tree.InOrderTreeTravalsal()
}

func (root *Node) PreOrderTreeTravalsal() {
	// fmt.Printf("Inside PreOrderTreeTravalsal: %d", root.Val)
	if root == nil {
		return
	}

	fmt.Printf("%d", root.Val)
	root.Left.PreOrderTreeTravalsal()
	root.Right.PreOrderTreeTravalsal()
	return
}

func (root *Node) PostOrderTreeTravalsal() {
	if root == nil {
		return
	}

	root.Left.PostOrderTreeTravalsal()
	root.Right.PostOrderTreeTravalsal()
	fmt.Printf("%d", root.Val)
	return
}

func (root *Node) InOrderTreeTravalsal() {
	if root == nil {
		return
	}

	root.Left.InOrderTreeTravalsal()
	fmt.Printf("%d", root.Val)
	root.Right.InOrderTreeTravalsal()

}
