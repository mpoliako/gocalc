package calculator

import (
	"bytes"
	"strings"
)

const (
	leaf = iota
	node = iota
)

type opTree struct {
	parent *opTree
	left   *opTree
	right  *opTree
	value  string
	typ    int
}

var (
	bracketMap map[string]string
)

func buildTree(expr string) *opTree {

	rootNode := &opTree{
		parent: nil,
		left:   nil,
		right:  nil,
		value:  expr,
		typ:    leaf,
	}
	rootNode = buildTreeBySeparator(rootNode, PLUS)
	rootNode = buildTreeBySeparator(rootNode, MINUS)
	rootNode = buildTreeBySeparator(rootNode, MULTIPLY)
	rootNode = buildTreeBySeparator(rootNode, DIVIDE)
	return rootNode
}

func buildTreeBySeparator(tree *opTree, separator string) *opTree {

	if tree.typ == leaf {
		if strings.Contains(tree.value, separator) {
			return buildLocalTree(tree.value, separator)
		} else {
			return tree
		}
	}

	if tree.left != nil {
		tree.left = buildTreeBySeparator(tree.left, separator)
	}
	if tree.right != nil {
		tree.right = buildTreeBySeparator(tree.right, separator)
	}

	return tree
}

func buildLocalTree(expr string, separator string) *opTree {
	rootNode := &opTree{
		parent: nil,
		left:   nil,
		right:  nil,
		value:  "",
		typ:    node,
	}
	tokens := strings.Split(expr, separator)
	for indx, leafExpr := range tokens {
		leafNode := opTree{
			parent: nil,
			left:   nil,
			right:  nil,
			value:  leafExpr,
			typ:    leaf,
		}
		rootNode.addLeaf(&leafNode)
		if indx != len(tokens)-1 {
			opNode := opTree{
				parent: nil,
				left:   nil,
				right:  nil,
				value:  separator,
				typ:    node,
			}
			rootNode = rootNode.addNode(&opNode)
		}
	}
	return rootNode
}

func (tree *opTree) addLeaf(node *opTree) bool {
	if tree.typ == leaf {
		return false
	}

	if tree.left != nil {
		if tree.left.addLeaf(node) {
			return true
		}
	} else {
		tree.left = node
		node.parent = tree
		return true
	}

	if tree.right != nil {
		if tree.right.addLeaf(node) {
			return true
		}
	} else {
		tree.right = node
		node.parent = tree
		return true
	}

	return false
}

func (tree *opTree) addNode(node *opTree) *opTree {
	if tree.value == "" {
		tree.value = node.value
		return tree
	}

	tree.parent = node
	node.left = tree
	return node
}

func (tree *opTree) String() string {
	var buff bytes.Buffer

	buff.WriteString(tree.value)

	if tree.left != nil {
		buff.WriteString(tree.left.String())
	}
	if tree.right != nil {
		buff.WriteString(tree.right.String())
	}
	return buff.String()
}
