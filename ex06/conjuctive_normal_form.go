package main

import (
	"fmt"
	"os"
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
	if len(*formula) == 0 {
		panic("formula is wrong")
	}
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

func (node *AstNode) IsIn(haystack string) bool {
	for _, c := range haystack {
		if node.Item[0] == byte(c) {
			return true
		}
	}
	return false
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

		node.LeftLeaf.LeftLeaf = leftCopy.Clone()

		node.NegationNormalForm()
	}
}

func (node *AstNode) Clone() *AstNode {
	// Create a deep clone of the node and its sub-tree
	if node == nil {
		return nil
	}
	clone := NewAstNode(node.Item)
	clone.LeftLeaf = node.LeftLeaf.Clone()
	clone.RightLeaf = node.RightLeaf.Clone()
	return clone
}

func (node *AstNode) Stringify() string {
	expr := ""
	if node.LeftLeaf != nil {
		expr += node.LeftLeaf.Stringify()
	}
	if node.RightLeaf != nil {
		expr += node.RightLeaf.Stringify()
	}
	expr += node.Item
	return expr
}

func (node *AstNode) IsConjunctive() bool {
	if node.Item == "|" {
		if (node.RightLeaf.Item == "&" || node.LeftLeaf.Item == "&") &&
			(node.RightLeaf.IsIn("&|") || node.LeftLeaf.IsIn("&|")) {
			return false
		}
	}
	return true
}

func (node *AstNode) ConjunctiveNormalForm() {
	if node.LeftLeaf != nil {
		node.LeftLeaf.ConjunctiveNormalForm()
	}

	if node.RightLeaf != nil {
		node.RightLeaf.ConjunctiveNormalForm()
	}

	if !node.IsConjunctive() {
		node.Item = "&"
		xCopy := node.LeftLeaf.Clone()
		aCopy := node.RightLeaf.LeftLeaf.Clone()
		bCopy := node.RightLeaf.RightLeaf.Clone()

		node.LeftLeaf = NewAstNode("|")
		node.RightLeaf = NewAstNode("|")

		node.LeftLeaf.LeftLeaf = xCopy
		node.LeftLeaf.RightLeaf = aCopy

		node.RightLeaf.LeftLeaf = xCopy
		node.RightLeaf.RightLeaf = bCopy
	}

	if strings.Contains("&|", node.Item) {
		if node.LeftLeaf != nil && node.LeftLeaf.Item == node.Item &&
			(node.RightLeaf.IsIn("!") || len(node.RightLeaf.Item) == 1) {
			rightChildCopy := node.RightLeaf.Clone()
			leftChildCopy := node.LeftLeaf.Clone()
			node.RightLeaf = leftChildCopy
			node.LeftLeaf = rightChildCopy
		}
	}
}

func isUpper(char rune) bool {
	return (char >= 'A' && char <= 'Z')
}

func reduceDoubleNegation(nnf string) string {
	for {
		i := 0
		for i < len(nnf) {
			if isUpper(rune(nnf[i])) && (i+1 < len(nnf) && nnf[i+1] == '!') &&
				(i+2 < len(nnf) && nnf[i+2] == '!') {
				nnf = nnf[:i+1] + nnf[i+3:]
				i = 0
			}
			i++
		}
		// break when string was not modified
		if i == len(nnf) {
			break
		}
	}
	return nnf
}

func ConjunctiveNormalForm(formula string) string {
	formulaStack := formula
	root := NewAstNode("0")
	root.ParseFormula(&formulaStack)
	root.NegationNormalForm()
	root.ConjunctiveNormalForm()
	cnf := root.Stringify()
	cnf = reduceDoubleNegation(cnf)
	return cnf
}

func isSpecial(c rune) bool {
	return (c == '&' || c == '!' || c == '^' || c == '>' || c == '=')
}

func isValid(s string) bool {
	for _, char := range s {
		if !isUpper(char) && !isSpecial(char) {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program 'logical formula'")
		return
	}
	if !isValid(os.Args[1]) {
		fmt.Println("Usage: program string_to_evaluate")
	}

	formula := os.Args[1]
	cnf := ConjunctiveNormalForm(formula)
	fmt.Println("str:", formula, "; CNF:", cnf)
	fmt.Println()

	formula = "AB&!"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB|!"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB|C&"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB|C|D|"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB&C&D&"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB&!C!|"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
	fmt.Println()

	formula = "AB|!C!&"
	fmt.Println("str:", formula, "; CNF:", ConjunctiveNormalForm(formula))
}
