package main

import (
	"fmt"
)

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	vis := make([]bool, n)
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return DFS(source, destination, vis, g)
}

func DFS(i, destination int, vis []bool, g [][]int) bool {
	if i == destination {
		return true
	}
	vis[i] = true
	for _, j := range g[i] {
		if !vis[j] && DFS(j, destination, vis, g) {
			return true
		}
	}
	return false
}

func main() {
	// [[0,1],[0,2],[3,5],[5,4],[4,3]]
	graph := [][]int{{0, 1}, {1, 2}, {3, 5}, {5, 4}, {4, 3}}
	isExist := ValidPath(6, graph, 0, 5)
	fmt.Println("isExist", isExist)
}
