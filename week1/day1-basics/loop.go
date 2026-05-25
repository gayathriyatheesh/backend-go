package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("Inner Loop Iteration:", j)
		}
		fmt.Println("Iteration:", i)
	}
}
