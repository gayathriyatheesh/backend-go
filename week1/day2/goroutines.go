package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(n int, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Println("%s, %d", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go printNumbers(5, fmt.Sprintf("Goroutine %d", i), &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}
