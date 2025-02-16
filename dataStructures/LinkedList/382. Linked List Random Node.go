package main

import (
	"time"

	"golang.org/x/exp/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		Head: head,
	}
}

func (this *Solution) GetRandom() int {
	idx := 0
	t := this.Head
	val := 0
	for t != nil {
		idx++
		rnd := rand.Intn(idx)
		if rnd == 0 {
			val = t.Val
		}
		t = t.Next
	}
	return val
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Example usage
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	solution := Constructor(head)

	// Get random values
	for i := 0; i < 5; i++ {
		println(solution.GetRandom())
	}
}
