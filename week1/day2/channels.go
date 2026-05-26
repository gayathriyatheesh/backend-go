package main

import (
	"fmt"
	"sync"
)

func sender(ch chan string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- msg
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go sender(ch, fmt.Sprintf("Message from sender %d", i), &wg)
	}

	for i := 1; i < 3; i++ {
		msg := <-ch
		fmt.Println("Received:", msg)
	}
	wg.Wait()
	fmt.Println("Done!")
}
