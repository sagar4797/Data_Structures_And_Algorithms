package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	// If the combined lengths of s1 and s2 do not match the length of s3, return false
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// Create a 2D dp array where dp[i][j] indicates if s1[0..i] and s2[0..j] form s3[0..i+j]
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// fmt.Println("dp: ", dp)
	// Base case: Empty s1 and s2 can form empty s3
	dp[len(s1)][len(s2)] = true

	// Fill the dp table
	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {

			if i < len(s1) && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}
			if j < len(s2) && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	// The answer is whether the full strings s1 and s2 can form s3
	return dp[0][0]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	isInterleaveVal1 := isInterleave(s1, s2, s3)
	fmt.Println("-------------->", isInterleaveVal1)

	s4 := "aabcc"
	s5 := "dbbca"
	s6 := "aadbbbaccc"
	isInterleaveVal2 := isInterleave(s4, s5, s6)
	fmt.Println("-------------->", isInterleaveVal2)
}
