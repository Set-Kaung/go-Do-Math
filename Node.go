package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func calculate(left, right int, op string) int {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			return 0
		}
		return left / right
	default:
		return 0
	}
}

// var operators = []string{"+", "-", "*", "/"}
func Evaluate(n *Node) int {
	if n.Left == nil && n.Right == nil {
		number, _ := strconv.Atoi(n.Data)
		return number
	}
	n.Left.Data = strconv.Itoa(Evaluate(n.Left))
	n.Right.Data = strconv.Itoa(Evaluate(n.Right))

	left, _ := strconv.Atoi(n.Left.Data)
	right, _ := strconv.Atoi(n.Right.Data)
	return calculate(left, right, n.Data)
}

func (n *Node) LNR() string {
	sb := &strings.Builder{}
	if n.Left != nil {
		sb.WriteString(n.Left.LNR())
	}
	sb.WriteString(n.Data)
	if n.Right != nil {
		sb.WriteString(n.Right.LNR())
	}
	return sb.String()
}

func (n *Node) NLR() string {
	sb := &strings.Builder{}
	sb.WriteString(n.Data)
	sb.WriteString(" ")
	if n.Left != nil {
		sb.WriteString(n.Left.LNR())
	}
	sb.WriteString(" ")
	if n.Right != nil {
		sb.WriteString(n.Right.LNR())
	}
	return sb.String()
}
