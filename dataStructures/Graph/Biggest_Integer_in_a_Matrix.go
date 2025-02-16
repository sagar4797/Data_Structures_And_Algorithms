
// Probelm link: https://leetcode.com/discuss/interview-question/645247/toptal-oa-2020-biggest-integer-in-a-matrix


// You are given a board of N rows and M columns. Each filed of the board contains a single digit (0-9)

// You want to find a path consisting of four neighboring fields. Two fields are neighboring if they share a common side. Also, the fields in your path should be distinct (you can't visit the same field twice).

// The four digits of your path, int the order in which you visit them, create an integer. what is the biggest integer that you can achieve in this way?

// Write a function

// func Solution(Board [][]int)int

// that, given the board represented as a matrix of integers consisting of N rows and M columns, returns the biggest integer that you can achieve when concatenating the values in a path of length four.
// Examples:
// Given the following board(N=3,M=5)
// [[9,1,1,0,7],[1,0,2,1,0],[1,9,1,1,0]]
// the function should return 9121. you can choose either of the following paths(the first field is denoted by red):
// Given the following board (N=3,M=3)
// [[1,1,1],[1,3,4],[1,4,3]]
// the function should return 4343. The best path is
// Given the following board(N=1, M=5)
// [0,1,5,0,0]
// should return 1500


package main

import (
	"fmt"
	"math"
	"strconv"
)

type Position struct {
	Row, Col int
}

var directions = [][]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func findMaxPathDFS(board [][]int, start Position, visited [][]bool, currentPath int) int {
	fmt.Println("start: ", start)
	visited[start.Row][start.Col] = true
	currentPath = currentPath*10 + board[start.Row][start.Col]
	fmt.Println("currentPath: ", currentPath)

	// Check if the path length is 4
	if len(strconv.Itoa(currentPath)) > 4 {
		fmt.Println("inside if currentPath: ", currentPath)
		visited[start.Row][start.Col] = false // Unmark the current position as visited
		return 0                              // Return 0 if the path length exceeds 4
	}

	fmt.Println("Continue..............")
	// Initialize max with the current path
	max := currentPath

	for _, dir := range directions {
		newRow := start.Row + dir[0]
		newCol := start.Col + dir[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && !visited[newRow][newCol] {
			fmt.Println("newRow: ", newRow, "newCol: ", newCol)
			fmt.Println("board", board)
			// Get the maximum path from the recursive call
			neighborMax := findMaxPathDFS(board, Position{newRow, newCol}, visited, currentPath)
			// Update max with the maximum of current max and neighborMax
			max = int(math.Max(float64(max), float64(neighborMax)))
		}
	}

	visited[start.Row][start.Col] = false // Unmark the current position as visited
	return max
}

func Solution(board [][]int) int {
	maxValue := 0
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Println("r: ", r, "c: ", c)
			maxValue = int(math.Max(float64(maxValue), float64(findMaxPathDFS(board, Position{r, c}, visited, 0))))
			fmt.Println("maxValue: ", maxValue)
		}
	}

	return maxValue
}

func main() {
	board1 := [][]int{
		{9, 1, 1, 0, 7},
		{1, 0, 2, 1, 0},
		{1, 9, 1, 1, 0},
	}
	fmt.Println(Solution(board1)) // Output: 9121

	board2 := [][]int{
		{1, 1, 1},
		{1, 3, 4},
		{1, 4, 3},
	}
	fmt.Println(Solution(board2)) // Output: 4343

	board3 := [][]int{
		{0, 1, 5, 0, 0},
	}
	fmt.Println(Solution(board3)) // Output: 1500
}
