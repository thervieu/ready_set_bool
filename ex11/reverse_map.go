package main

import "fmt"

func reverse_map(mapped float64) (uint16, uint16) {
	combination := int32(mapped * float64(0xFFFF_FFFF))

	return uint16(combination>>16), uint16(combination & 0xFFFF)
}

func main() {
	mapped := float64(0)
	x, y := reverse_map(mapped)
	fmt.Println(mapped, "=>", x, "and", y)

	mapped = float64(2.3283064370807974e-10)
	x, y = reverse_map(mapped)
	fmt.Println(mapped, "=>", x, "and", y)

	mapped = float64(1.5258789066052714e-05)
	x, y = reverse_map(mapped)
	fmt.Println(mapped, "=>", x, "and", y)

	mapped = float64(0.00015258882198310197)
	x, y = reverse_map(mapped)
	fmt.Println(mapped, "=>", x, "and", y)

	mapped = float64(0.06248480059730001)
	x, y = reverse_map(mapped)
	fmt.Println(mapped, "=>", x, "and", y)
}