package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		flag := true
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			if stack[len(stack)-1] == -asteroid {
				fmt.Println("before stack= ", stack)
				fmt.Println("stack[:len(stack)-1]= ", stack[:len(stack)-1])
				stack = stack[:len(stack)-1]
				fmt.Println("after stack= ", stack)
			} else if stack[len(stack)-1] < -asteroid {
				fmt.Println("stack[len(stack)-1]==== ", stack[len(stack)-1])
				fmt.Println("Inside -asteroid", asteroid)
				stack = stack[:len(stack)-1]
				fmt.Println("stack@@@@@@@@", stack)
				continue
			}
			flag = false
			break
		}

		if flag {
			fmt.Println("asteroid= ", asteroid)
			stack = append(stack, asteroid)
		}
	}
	return stack
}

func main() {
	asteroids := []int{4, 7, 1, 1, 2, -3, -7, 17, 15, -16}
	outPut := asteroidCollision(asteroids)
	fmt.Println("outPut:", outPut)
}
