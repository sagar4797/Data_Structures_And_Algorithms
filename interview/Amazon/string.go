package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := "sagar"
	// for _, char := range str {
	// 	if char == 'g' {
	// 		fmt.Println(string(char))
	// 	}

	// }
	s := "abcde"
	goal := "cdeab"
	rotateString(s, goal)
}

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}

	arr := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		arr = append(arr[1:], arr[:1]...)
		if goal == strings.Join(arr, "") {
			return true
		}
	}

	return false
}
