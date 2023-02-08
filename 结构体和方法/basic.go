package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node treeNode) getValue() int {
	return node.value
}
func (node *treeNode) setValue(value int) {
	node.value = value
}
func main() {
	var root treeNode
	root.value = 1
	root.left = &treeNode{}
	root.right = &treeNode{}
	root.right.left = &treeNode{2, nil, nil}
	fmt.Println(root.getValue())
	root.setValue(4)
	fmt.Println(root.getValue())
}
