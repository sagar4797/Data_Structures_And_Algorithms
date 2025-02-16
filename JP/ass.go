package main

import (
	"fmt"
	"math"
	"strconv"
)

type Position struct {
	Row int
	Col int
}

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func findMapPathDFS(board [][]int, start Position, visited [][]bool, currentPath int) int {
	visited[start.Row][start.Col] = true
	println("currentPath--before", currentPath)
	currentPath = currentPath*10 + board[start.Row][start.Col]
	println("currentPath--after", currentPath)
	if len(strconv.Itoa(currentPath)) > 4 {
		visited[start.Row][start.Col] = false
		return 0
	}

	max := currentPath

	for _, dir := range directions {
		newRow := start.Row + dir[0]
		newCol := start.Col + dir[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && !visited[newRow][newCol] {
			neighboreMax := findMapPathDFS(board, Position{newRow, newCol}, visited, currentPath)
			max = int(math.Max(float64(max), float64(neighboreMax)))
		}
	}

	visited[start.Row][start.Col] = false
	return max
}

func Solution(board [][]int) int {
	// Implement your solution here

	maxValue := 0
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			maxValue = int(math.Max(float64(maxValue), float64(findMapPathDFS(board, Position{r, c}, visited, 0))))
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
	Spath1 := Solution(board1)

	fmt.Println("Spath1", Spath1)
}
