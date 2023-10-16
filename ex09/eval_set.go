package main

import (
	"fmt"
	"sort"
	"strings"
)

type AstNode struct {
	Item      string
	LeftLeaf  *AstNode
	RightLeaf *AstNode
}

func NewAstNode(item string) *AstNode {
	return &AstNode{
		Item:      item,
		LeftLeaf:  nil,
		RightLeaf: nil,
	}
}

func (node *AstNode) ParseFormula(formula *string) {
	operands := "!&|^>="
	node.Item = string((*formula)[len(*formula)-1])
	*formula = (*formula)[:len(*formula)-1]

	if strings.ContainsRune(operands, rune(node.Item[0])) {
		if node.Item != "!" {
			node.RightLeaf = NewAstNode("0")
			node.RightLeaf.ParseFormula(formula)
		}
		node.LeftLeaf = NewAstNode("0")
		node.LeftLeaf.ParseFormula(formula)
	}
}

func bytesContain(slice []byte, item byte) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (node *AstNode) IsIn(haystack string) bool {
	for _, c := range haystack {
		if node.Item[0] == byte(c) {
			return true
		}
	}
	return false
}

func (node *AstNode) Clone() *AstNode {
	if node == nil {
		return nil
	}
	clone := NewAstNode(node.Item)
	clone.LeftLeaf = node.LeftLeaf.Clone()
	clone.RightLeaf = node.RightLeaf.Clone()
	return clone
}

func (node *AstNode) NegationNormalForm() {
	if node.LeftLeaf != nil {
		node.LeftLeaf.NegationNormalForm()
	}

	if node.RightLeaf != nil {
		node.RightLeaf.NegationNormalForm()
	}

	if node.Item == "!" && node.LeftLeaf.IsIn("&|") {
		rightCopy := node.LeftLeaf.RightLeaf
		if node.LeftLeaf.Item == "|" {
			node.Item = "&"
		} else {
			node.Item = "|"
		}

		node.LeftLeaf.Item = "!"
		node.LeftLeaf.RightLeaf = nil

		node.RightLeaf = NewAstNode("!")
		node.RightLeaf.LeftLeaf = rightCopy

		node.NegationNormalForm()
	}

	if node.Item == "=" {
		node.Item = "&"
		aCopy := node.LeftLeaf
		bCopy := node.RightLeaf

		node.LeftLeaf = NewAstNode(">")
		node.RightLeaf = NewAstNode(">")

		node.LeftLeaf.LeftLeaf = aCopy.Clone()
		node.LeftLeaf.RightLeaf = bCopy.Clone()

		node.RightLeaf.LeftLeaf = bCopy.Clone()
		node.RightLeaf.RightLeaf = aCopy.Clone()

		node.NegationNormalForm()
	}

	if node.Item == "^" {
		node.Item = "|"
		aCopy := node.LeftLeaf
		bCopy := node.RightLeaf

		node.LeftLeaf = NewAstNode("&")
		node.RightLeaf = NewAstNode("&")

		node.LeftLeaf.RightLeaf = NewAstNode("!")
		node.LeftLeaf.RightLeaf.LeftLeaf = bCopy.Clone()
		node.LeftLeaf.LeftLeaf = aCopy.Clone()

		node.RightLeaf.LeftLeaf = NewAstNode("!")
		node.RightLeaf.LeftLeaf.LeftLeaf = aCopy.Clone()
		node.RightLeaf.RightLeaf = bCopy.Clone()

		node.NegationNormalForm()
	}

	if node.Item == ">" {
		leftCopy := node.LeftLeaf
		node.Item = "|"

		node.LeftLeaf = NewAstNode("!")
		node.LeftLeaf.LeftLeaf = leftCopy

		node.NegationNormalForm()
	}
}

func (node *AstNode) Compute(sets [][]int, superset []int) []int {
	var result []int

	if !strings.Contains("&|!", string(node.Item)) {
		result = sets[int(byte(node.Item[0])-'A')]
	} else if node.Item == "!" {
		for _, element := range superset {
			if !intSliceContains(node.LeftLeaf.Compute(sets, superset), element) {
				result = append(result, element)
			}
		}
	} else if node.Item == "|" {
		set := make(map[int]struct{})
		setSlice(set, node.LeftLeaf.Compute(sets, superset))
		setSlice(set, node.RightLeaf.Compute(sets, superset))
		result = setToSlice(set)
	} else if node.Item == "&" {
		for _, x := range node.LeftLeaf.Compute(sets, superset) {
			if intSliceContains(node.RightLeaf.Compute(sets, superset), x) {
				result = append(result, x)
			}
		}
	}
	return result
}

func intSliceContains(slice []int, item int) bool {
	for _, x := range slice {
		if x == item {
			return true
		}
	}
	return false
}

func setSlice(set map[int]struct{}, slice []int) {
	for _, x := range slice {
		set[x] = struct{}{}
	}
}

func setToSlice(set map[int]struct{}) []int {
	var slice []int
	for x := range set {
		slice = append(slice, x)
	}
	sort.Ints(slice)
	return slice
}

func IsValidFormula(formula string, setsSize int) bool {
	charSet := make(map[rune]struct{})
	for _, c := range formula {
		if c >= 'A' && c <= 'Z' {
			charSet[c] = struct{}{}
		}
	}
	return len(charSet) == setsSize
}

func GetSuperset(sets [][]int) []int {
	set := make(map[int]struct{})
	for _, elements := range sets {
		setSlice(set, elements)
	}
	return setToSlice(set)
}

func EvalSet(formula string, sets [][]int) []int {
	if !IsValidFormula(formula, len(sets)) {
		panic("The formula and sets provided are not compatible")
	}

	formulaStack := string(formula)
	root := NewAstNode("0")
	root.ParseFormula(&formulaStack)
	root.NegationNormalForm()
	return root.Compute(sets, GetSuperset(sets))
}

func main() {
	sets := [][]int{
		{0, 2, 3},
		{0, 4, 5},
	}
	formula := "AB&"
	result := EvalSet(formula, sets)
	fmt.Println("formula", formula, "\nsets", sets, "\nResulting Set:", result)
	fmt.Println()

	sets = [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	formula = "AB|"
	result = EvalSet(formula, sets)
	fmt.Println("formula", formula, "\nsets", sets, "\nResulting Set:", result)
	fmt.Println()

	sets = [][]int{
		{0, 1, 2},
	}
	formula = "A!"
	result = EvalSet(formula, sets)
	fmt.Println("formula", formula, "\nsets", sets, "\nResulting Set:", result)
}
