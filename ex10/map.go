package main

import "fmt"

func map2(x uint16, y uint16) float64 {
	combination := uint32(x)<<16 + uint32(y)
	mapped := float64(combination) / float64(0xFFFF_FFFF)

	return mapped
}

func main() {
	x := uint16(0)
	y := uint16(0)
	fmt.Println(x, "and", y, "=>", map2(x, y))

	x = uint16(0)
	y = uint16(1)
	fmt.Println(x, "and", y, "=>", map2(x, y))

	x = uint16(1)
	y = uint16(0)
	fmt.Println(x, "and", y, "=>", map2(x, y))

	x = uint16(10)
	y = uint16(4)
	fmt.Println(x, "and", y, "=>", map2(x, y))

	x = uint16(0xFFF)
	y = uint16(0x0FF)
	fmt.Println(x, "and", y, "=>", map2(x, y))
}
