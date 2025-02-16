package main

import (
	"fmt"
	"sync"
)

type Person struct {
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(1)
	mu.Lock()
	err := Producer(ch, &wg)
	if err != nil {

	}
	mu.Unlock()
	wg.Add(1)
	go Consumer(ch, &wg)

	wg.Wait()
}

func Producer(ch chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}
