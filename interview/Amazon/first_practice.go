package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func processLogs(logs []string, threshold int32) []string {
	// Write your code here
	userCount := make(map[int]int)
	// userArray := []string{}

	for _, log := range logs {
		parts := strings.Split(log, " ")
		var transaction []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				transaction = append(transaction, num)
			}
		}
		userCount[transaction[0]]++
		userCount[transaction[1]]++
	}

	var outArray []int
	fmt.Println("userCountArray: %+v", userCount)
	for n, c := range userCount {
		if int(threshold) <= c {
			fmt.Println(n, c)
			outArray = append(outArray, n)
			fmt.Println("outString = ", outArray)
		}
	}

	fmt.Println("OutArray: ", outArray)

	sort.Ints(outArray)
	outString := []string{}
	for _, n := range outArray {
		outString = append(outString, strconv.Itoa(n))
	}

	return outString
}

func main() {
	logs := []string{"1 2 50", "1 7 70", "1 3 20", "2 2 17"}
	out := processLogs(logs, 2)
	fmt.Println("OutPut", out)
}
