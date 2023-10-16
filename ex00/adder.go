package main

import "fmt"

func adder(a uint32, b uint32) uint32 {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	// 0 + 0
	fmt.Printf("%d + %d = %d\n", uint32(0), uint32(0), adder(uint32(0), uint32(0)))
	// 1 + 0
	fmt.Printf("%d + %d = %d\n", uint32(1), uint32(0), adder(uint32(1), uint32(0)))
	// 0 + 1
	fmt.Printf("%d + %d = %d\n", uint32(0), uint32(1), adder(uint32(0), uint32(1)))
	// 1 + 1
	fmt.Printf("%d + %d = %d\n", uint32(1), uint32(1), adder(uint32(1), uint32(1)))
	// 2 + 1
	fmt.Printf("%d + %d = %d\n", uint32(2), uint32(1), adder(uint32(2), uint32(1)))
	// 2 + 2
	fmt.Printf("%d + %d = %d\n", uint32(2), uint32(2), adder(uint32(2), uint32(2)))
	// 7 + 7
	fmt.Printf("%d + %d = %d\n", uint32(7), uint32(7), adder(uint32(7), uint32(7)))
	// 123 + 123
	fmt.Printf("%d + %d = %d\n", uint32(123), uint32(123), adder(uint32(123), uint32(123)))
	// 4294967295 + 4294967295
	fmt.Printf("%d + %d = %d\n", uint32(4294967295), uint32(4294967295), adder(uint32(4294967295), uint32(4294967295)))
}
