package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go Producer(ch, &wg)
	wg.Add(1)
	go Reciver(ch, &wg)
	wg.Wait()
	mu.Lock()
	mu.Unlock()
}

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func Reciver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}
