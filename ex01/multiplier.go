package main

import "fmt"

func multiplier(a uint32, b uint32) uint32 {
	var result uint32 = 0

	for b > 0 {
		if b&1 == 1 {
			result += a
		}

		a <<= 1
		b >>= 1
	}
	return result
}

func main() {
	// 0 * 0
	fmt.Printf("%d * %d = %d\n", uint32(0), uint32(0), multiplier(uint32(0), uint32(0)))
	// 1 * 0
	fmt.Printf("%d * %d = %d\n", uint32(1), uint32(0), multiplier(uint32(1), uint32(0)))
	// 0 * 1
	fmt.Printf("%d * %d = %d\n", uint32(0), uint32(1), multiplier(uint32(0), uint32(1)))
	// 1 * 1
	fmt.Printf("%d * %d = %d\n", uint32(1), uint32(1), multiplier(uint32(1), uint32(1)))
	// 2 * 1
	fmt.Printf("%d * %d = %d\n", uint32(2), uint32(1), multiplier(uint32(2), uint32(1)))
	// 2 * 2
	fmt.Printf("%d * %d = %d\n", uint32(2), uint32(2), multiplier(uint32(2), uint32(2)))
	// 7 * 7
	fmt.Printf("%d * %d = %d\n", uint32(7), uint32(7), multiplier(uint32(7), uint32(7)))
	// 9 * 10
	fmt.Printf("%d * %d = %d\n", uint32(9), uint32(10), multiplier(uint32(9), uint32(10)))
	// 5 * 25
	fmt.Printf("%d * %d = %d\n", uint32(5), uint32(25), multiplier(uint32(5), uint32(25)))
	// 123 * 123
	fmt.Printf("%d * %d = %d | truth: %d\n", uint32(123), uint32(123), multiplier(uint32(123), uint32(123)), 123*123)
	// 4294967295 * 4294967295
	fmt.Printf("%d * %d = %d\n", uint32(4294967295), uint32(4294967295), multiplier(uint32(4294967295), uint32(4294967295)))
}
