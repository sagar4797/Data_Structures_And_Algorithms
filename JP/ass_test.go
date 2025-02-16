package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		board [][]int
		want  int
	}{
		{
			board: [][]int{
				{9, 1, 1, 0, 7},
				{1, 0, 2, 1, 0},
				{1, 9, 1, 1, 0},
			},
			want: 9121,
		},
		{
			board: [][]int{
				{1, 1, 1},
				{1, 3, 4},
				{1, 4, 3},
			},
			want: 4343,
		},
		{
			board: [][]int{
				{0, 1, 5, 0, 0},
			},
			want: 1500,
		},
		{
			board: [][]int{
				{5, 5, 5, 5},
				{5, 5, 5, 5},
				{5, 5, 5, 5},
			},
			want: 5555,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Solution(tt.board); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}

}
