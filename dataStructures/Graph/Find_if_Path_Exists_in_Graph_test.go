package main

import (
	"fmt"
	"testing"
)

func TestValidPath(t *testing.T) {
	tests := []struct {
		n      int
		edges  [][]int
		source int
		dest   int
		want   bool
	}{
		{6, [][]int{{0, 1}, {1, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
	}
	for _, test := range tests {
		actual := ValidPath(test.n, test.edges, test.source, test.dest)
		if actual != test.want {
			t.Errorf("Test case failed! Expected: %v, Got: %v", test.want, actual)
		}
	}

}

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")
	// Call the test function
	TestValidPath(m)
}
