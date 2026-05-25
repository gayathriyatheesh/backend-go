package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a := 10
	b := 20

	fmt.Println("Before Swap: x =", a, "y =", b)
	a, b = swap(a, b)
	fmt.Println("After Swap: x =", a, "y =", b)
}
