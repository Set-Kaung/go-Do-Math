package main

import "fmt"

var operatorsLevel = map[rune]int{'-': 4, '+': 3, '/': 2, '*': 1}

func containsOperator(expression string) bool {
	for _, r := range expression {
		if _, ok := operatorsLevel[r]; ok {
			return ok
		}
	}
	return false
}

func splitRecurse(expression string) *Node {

	highest := '*'
	index := 0
	for i, r := range expression {
		if level, ok := operatorsLevel[r]; ok {
			if level > operatorsLevel[highest] {
				highest = r
				index = i
			} else if level == operatorsLevel[highest] && i > 0 {
				index = i
			}
		}
	}
	left, right := expression[:index], expression[index+1:]
	t := &Node{}
	t.Data = string(highest)
	if containsOperator(left) {
		t.Left = splitRecurse(left)
	} else {
		t.Left = &Node{Data: left}
	}
	if containsOperator(right) {
		t.Right = splitRecurse(right)
	} else {
		t.Right = &Node{Data: right}
	}

	return t
}

func main() {
	express := "2+2"
	tree := splitRecurse(express)
	fmt.Println(tree.LNR())
	fmt.Println(tree.NLR())
}
