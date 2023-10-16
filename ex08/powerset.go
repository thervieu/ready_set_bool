package main

import "fmt"

func powerset(set []int) [][]int {
	var powerset [][]int
	powerset = append(powerset, []int{})

	for _, elt := range set {
		var subsets [][]int
		for _, subset := range powerset {
			newSubset := append([]int{elt}, subset...)
			subsets = append(subsets, newSubset)
		}
		powerset = append(powerset, subsets...)
	}
	return powerset
}

func main() {
	set1 := []int{1}
	fmt.Println("Powerset of ", set1, "is", powerset(set1))
	set2 := []int{1, 2}
	fmt.Println("Powerset of ", set2, "is", powerset(set2))
	set3 := []int{1, 2, 3}
	fmt.Println("Powerset of ", set3, "is", powerset(set3))
	return
}