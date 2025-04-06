package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Test input with invalid tree
	treeString2 := []string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"}

	rNode12 := buildBinaryTree(treeString2)
	isValidBTre2 := validateBinaryTree(rNode12)
	println("is valid binary tree: ", isValidBTre2)
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func buildBinaryTree(treeS []string) *Node {
	tree := make(map[int][]int)
	childParentMap := make(map[int]int) // To track if a node has more than one parent

	// Parse the tree strings to build parent-child relationships
	for _, part := range treeS {
		// Remove parentheses and split the string
		trimmed := strings.Trim(part, "()")
		nodes := strings.Split(trimmed, ",")
		if len(nodes) != 2 {
			fmt.Println("Invalid format")
			continue
		}

		// Convert strings to integers
		parent, err1 := strconv.Atoi(nodes[0])
		child, err2 := strconv.Atoi(nodes[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers")
			continue
		}

		// Check if the child already has a parent
		if existingParent, exists := childParentMap[child]; exists {
			// If the child has an existing parent, it violates the binary tree rules
			fmt.Println("Invalid binary tree: Node", child, "has multiple parents:", existingParent, "and", parent)
			return nil
		}

		// Assign the child to the parent
		childParentMap[child] = parent

		// Add the child to the parent list
		tree[parent] = append(tree[parent], child)
	}

	// Build the actual tree from the relationships
	nodeMap := make(map[int]*Node)
	var root *Node

	// Create nodes and link them as per the parent-child relationships
	for parent, children := range tree {
		// Create or get the parent node
		if _, exists := nodeMap[parent]; !exists {
			nodeMap[parent] = &Node{Val: parent}
		}

		// Check and assign children
		for _, child := range children {
			// Create or get the child node
			if _, exists := nodeMap[child]; !exists {
				nodeMap[child] = &Node{Val: child}
			}

			// Check where to place the child: left or right
			if nodeMap[parent].left == nil {
				nodeMap[parent].left = nodeMap[child]
			} else if nodeMap[parent].right == nil {
				nodeMap[parent].right = nodeMap[child]
			} else {
				// If the parent already has two children, it's invalid
				fmt.Println("Invalid binary tree: More than two children for parent", parent)
				return nil
			}
		}
	}

	// Find the root (the node that does not appear as a child)
	for _, node := range nodeMap {
		if _, exists := childParentMap[node.Val]; !exists {
			root = node
			break
		}
	}

	// If we cannot find the root (i.e., there's no node with no parent), it's invalid
	if root == nil {
		fmt.Println("Invalid binary tree: No root found.")
		return nil
	}

	return root
}

// Function to validate the binary tree
func validateBinaryTree(root *Node) bool {
	// We will perform a simple validation:
	// A valid binary tree has at most two children per node.
	// For simplicity, we use a DFS or BFS traversal to ensure no node has more than two children.

	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == nil {
			return true
		}

		// Count non-nil children (left and right)
		childrenCount := 0
		if node.left != nil {
			childrenCount++
		}
		if node.right != nil {
			childrenCount++package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Test input with invalid tree
	treeString2 := []string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"}

	rNode12 := buildBinaryTree(treeString2)
	isValidBTre2 := validateBinaryTree(rNode12)
	println("is valid binary tree: ", isValidBTre2)
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func buildBinaryTree(treeS []string) *Node {
	tree := make(map[int][]int)
	childParentMap := make(map[int]int) // To track if a node has more than one parent

	// Parse the tree strings to build parent-child relationships
	for _, part := range treeS {
		// Remove parentheses and split the string
		trimmed := strings.Trim(part, "()")
		nodes := strings.Split(trimmed, ",")
		if len(nodes) != 2 {
			fmt.Println("Invalid format")
			continue
		}

		// Convert strings to integers
		parent, err1 := strconv.Atoi(nodes[0])
		child, err2 := strconv.Atoi(nodes[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers")
			continue
		}

		// Check if the child already has a parent
		if existingParent, exists := childParentMap[child]; exists {
			// If the child has an existing parent, it violates the binary tree rules
			fmt.Println("Invalid binary tree: Node", child, "has multiple parents:", existingParent, "and", parent)
			return nil
		}

		// Assign the child to the parent
		childParentMap[child] = parent

		// Add the child to the parent list
		tree[parent] = append(tree[parent], child)
	}

	// Build the actual tree from the relationships
	nodeMap := make(map[int]*Node)
	var root *Node

	// Create nodes and link them as per the parent-child relationships
	for parent, children := range tree {
		// Create or get the parent node
		if _, exists := nodeMap[parent]; !exists {
			nodeMap[parent] = &Node{Val: parent}
		}

		// Check and assign children
		for _, child := range children {
			// Create or get the child node
			if _, exists := nodeMap[child]; !exists {
				nodeMap[child] = &Node{Val: child}
			}

			// Check where to place the child: left or right
			if nodeMap[parent].left == nil {
				nodeMap[parent].left = nodeMap[child]
			} else if nodeMap[parent].right == nil {
				nodeMap[parent].right = nodeMap[child]
			} else {
				// If the parent already has two children, it's invalid
				fmt.Println("Invalid binary tree: More than two children for parent", parent)
				return nil
			}
		}
	}

	// Find the root (the node that does not appear as a child)
	for _, node := range nodeMap {
		if _, exists := childParentMap[node.Val]; !exists {
			root = node
			break
		}
	}

	// If we cannot find the root (i.e., there's no node with no parent), it's invalid
	if root == nil {
		fmt.Println("Invalid binary tree: No root found.")
		return nil
	}

	return root
}

// Function to validate the binary tree
func validateBinaryTree(root *Node) bool {
	// We will perform a simple validation:
	// A valid binary tree has at most two children per node.
	// For simplicity, we use a DFS or BFS traversal to ensure no node has more than two children.

	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == nil {
			return true
		}

		// Count non-nil children (left and right)
		childrenCount := 0
		if node.left != nil {
			childrenCount++
		}
		if node.right != nil {
			childrenCount++
		}

		// If a node has more than two children, it's invalid
		if childrenCount > 2 {
			return false
		}

		// Recursively validate the left and right subtrees
		return dfs(node.left) && dfs(node.right)
	}

	return dfs(root)
}

		}

		// If a node has more than two children, it's invalid
		if childrenCount > 2 {
			return false
		}

		// Recursively validate the left and right subtrees
		return dfs(node.left) && dfs(node.right)
	}

	return dfs(root)
}




package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Test input with invalid tree
	treeString2 := []string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"}

	rNode12 := buildBinaryTree(treeString2)
	isValidBTre2 := validateBinaryTree(rNode12)
	println("is valid binary tree: ", isValidBTre2)
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func buildBinaryTree(treeS []string) *Node {
	tree := make(map[int][]int)
	childParentMap := make(map[int]int) // To track if a node has more than one parent

	// Parse the tree strings to build parent-child relationships
	for _, part := range treeS {
		// Remove parentheses and split the string
		trimmed := strings.Trim(part, "()")
		nodes := strings.Split(trimmed, ",")
		if len(nodes) != 2 {
			fmt.Println("Invalid format")
			continue
		}

		// Convert strings to integers
		parent, err1 := strconv.Atoi(nodes[0])
		child, err2 := strconv.Atoi(nodes[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers")
			continue
		}

		// Check if the child already has a parent
		if existingParent, exists := childParentMap[child]; exists {
			// If the child has an existing parent, it violates the binary tree rules
			fmt.Println("Invalid binary tree: Node", child, "has multiple parents:", existingParent, "and", parent)
			return nil
		}

		// Assign the child to the parent
		childParentMap[child] = parent

		// Add the child to the parent list
		tree[parent] = append(tree[parent], child)
	}

	// Build the actual tree from the relationships
	nodeMap := make(map[int]*Node)
	var root *Node

	// Create nodes and link them as per the parent-child relationships
	for parent, children := range tree {
		// Create or get the parent node
		if _, exists := nodeMap[parent]; !exists {
			nodeMap[parent] = &Node{Val: parent}
		}

		// Check and assign children
		for _, child := range children {
			// Create or get the child node
			if _, exists := nodeMap[child]; !exists {
				nodeMap[child] = &Node{Val: child}
			}

			// Check where to place the child: left or right
			if nodeMap[parent].left == nil {
				nodeMap[parent].left = nodeMap[child]
			} else if nodeMap[parent].right == nil {
				nodeMap[parent].right = nodeMap[child]
			} else {
				// If the parent already has two children, it's invalid
				fmt.Println("Invalid binary tree: More than two children for parent", parent)
				return nil
			}
		}
	}

	// Find the root (the node that does not appear as a child)
	for _, node := range nodeMap {
		if _, exists := childParentMap[node.Val]; !exists {
			root = node
			break
		}
	}

	// If we cannot find the root (i.e., there's no node with no parent), it's invalid
	if root == nil {
		fmt.Println("Invalid binary tree: No root found.")
		return nil
	}

	return root
}

// Function to validate the binary tree
func validateBinaryTree(root *Node) bool {
	// We will perform a simple validation:
	// A valid binary tree has at most two children per node.
	// For simplicity, we use a DFS or BFS traversal to ensure no node has more than two children.

	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == nil {
			return true
		}

		// Count non-nil children (left and right)
		childrenCount := 0
		if node.left != nil {
			childrenCount++
		}
		if node.right != nil {
			childrenCount++
		}

		// If a node has more than two children, it's invalid
		if childrenCount > 2 {
			return false
		}

		// Recursively validate the left and right subtrees
		return dfs(node.left) && dfs(node.right)
	}

	return dfs(root)
}
