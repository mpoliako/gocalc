// calculator project calculator.go
package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
)

func NewOp(op string) func(a int, b int) float64 {
	switch op {
	case PLUS:
		return func(a int, b int) float64 {
			return a + b
		}
	case MINUS:
		return func(a int, b int) float64 {
			return a - b
		}
	case MULTIPLY:
		return func(a int, b int) float64 {
			return a * b
		}
	case DIVIDE:
		return func(a int, b int) float64 {
			return a / b
		}
	default:
		panic(fmt.Sprintf("Not such function %v", op))
	}
}

func Calculate(expr string) float64 {
	// remove spaces
	expr = strings.Replace(expr, " ", "", -1)
	// build tree
	opTree := buildTree(expr)
	// calculate
	return calculate(opTree)
}

func calculate(tree *opTree) float64 {
	if tree.typ == leaf {
		ret, err := strconv.ParseFloat(tree.value, 64)
		if err != nil {
			panic(err)
		} else {
			return ret
		}
	}
	leftRes := calculate(tree.left)
	rightRes := calculate(tree.right)
	switch tree.value {
	case PLUS:
		return leftRes + rightRes
	case MINUS:
		return leftRes - rightRes
	case MULTIPLY:
		return leftRes * rightRes
	case DIVIDE:
		return leftRes / rightRes
	default:
		panic(fmt.Sprintf("Invalid operation %v", tree.value))
	}
}
