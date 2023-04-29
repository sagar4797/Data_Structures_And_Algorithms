package main

import "fmt"

func main() {
	arr := []int{7, 2, 12, 8, 3}
	fmt.Println("arrr: ", arr)

	len := len(arr)
	st := len - 1
	for i := 0; i < len-1; i++ {
		for j := 0; j < st; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		st--
	}

	fmt.Println("This is sorted arr: ", arr)
}
