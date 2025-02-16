package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	j := 0
	k := 0
	i := 0
	for i < len(s3) && j < len(s1) && k < len(s2) {
		if s3[i] == s2[k] {
			i++
			k++
			continue
		} else if s3[i] == s1[j] {
			i++
			j++
			continue
		} else {
			return false
		}
	}

	for i < len(s3) && j < len(s1) {
		if s3[i] == s1[j] {
			i++
			j++
			continue
		} else {
			return false
		}
	}

	for i < len(s3) && j < len(s2) {
		if s3[i] == s2[k] {
			fmt.Println("s3[i]", string(s3[i]))
			fmt.Println("s2[k]", string(s2[k]))
			i++
			k++
			continue
		} else {
			return false
		}
	}

	return true
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
