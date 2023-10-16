package main

import "fmt"

func gray(n uint32) uint32 {
	return n ^ (n >> 1)
}

func main() {
	fmt.Printf("%d => %d\n", 0, gray(0))
	fmt.Printf("%d => %d\n", 1, gray(1))
	fmt.Printf("%d => %d\n", 2, gray(2))
	fmt.Printf("%d => %d\n", 3, gray(3))
	fmt.Printf("%d => %d\n", 4, gray(4))
	fmt.Printf("%d => %d\n", 5, gray(5))
	fmt.Printf("%d => %d\n", 6, gray(6))
	fmt.Printf("%d => %d\n", 7, gray(7))
	fmt.Printf("%d => %d\n", 8, gray(8))
}
