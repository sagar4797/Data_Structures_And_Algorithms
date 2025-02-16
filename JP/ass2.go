package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Solution1(message string, k int) string {
// 	words := strings.Split(message, " ")
// 	if len(message) <= k {
// 		return message
// 	}

// 	endDots := "..."
// 	endDotsLen := len(endDots) + 1
// 	persistLen := k - endDotsLen

// 	var result []string
// 	currLen := 0

// 	for _, word := range words {
// 		wordLen := len(word)
// 		if currLen+wordLen > persistLen {
// 			break
// 		}
// 		result = append(result, word)
// 		currLen += wordLen + 1
// 	}

// 	if len(result) == 0 {
// 		return endDots
// 	}

// 	return strings.Join(result, " ") + " " + endDots
// }

// func main() {
// 	message1 := "And now here is my secret"
// 	fmt.Println(Solution1(message1, 15))
// }
