package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(s, t)
	fmt.Println("Result: ", result)
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Step 1: Count frequency of characters in t
	tCount := make(map[rune]int)
	for _, char := range t {
		tCount[char]++
	}

	// Step 2: Initialize the sliding window
	left, right := 0, 0
	required := len(tCount)
	formed := 0
	windowCount := make(map[rune]int)

	minLen := math.MaxInt32
	result := ""

	// Step 3: Start the sliding window
	for right < len(s) {
		// Add character from the right to the window
		char := rune(s[right])
		windowCount[char]++

		if count, exists := tCount[char]; exists && windowCount[char] == count {
			formed++
		}

		// Step 4: Try to contract the window until it's invalid
		for left <= right && formed == required {
			// Update the minimum window if necessary
			if right-left+1 < minLen {
				minLen = right - left + 1
				result = s[left : right+1]
			}

			// Remove character from the left
			leftChar := rune(s[left])
			windowCount[leftChar]--

			if count, exists := tCount[leftChar]; exists && windowCount[leftChar] < count {
				formed--
			}

			left++
		}

		// Expand the window from the right
		right++
	}

	return result
}
