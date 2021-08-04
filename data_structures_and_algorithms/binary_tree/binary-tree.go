package main

import (
	"fmt"
)

type TreeNode struct {
	right *TreeNode
	left  *TreeNode
	value int
}

type BinaryTree struct {
	head *TreeNode
}

func (tree *BinaryTree) Search(arg int) *TreeNode {
	return tree.search(arg, tree.head)
}

func (tree *BinaryTree) search(arg int, node *TreeNode) *TreeNode {
	switch {
	case node.value == arg:
		return node
	case arg < node.value:
		return tree.search(arg, node.left)
	case arg > node.value:
		return tree.search(arg, node.right)
	default:
		return nil
	}
}

func (tree *BinaryTree) Insert(arg int) {
	tree.insert(arg, tree.head)
}

func (tree *BinaryTree) insert(arg int, node *TreeNode) {
	if arg < node.value {
		if node.left == nil {
			node.left = &TreeNode{left: nil, right: nil, value: arg}
		} else {
			tree.insert(arg, node.left)
		}
	} else if arg > node.value {
		if node.right == nil {
			node.right = &TreeNode{left: nil, right: nil, value: arg}
		} else {
			tree.insert(arg, node.right)
		}
	}
}

func (tree *BinaryTree) Delete(arg int) {
	tree.delete(arg, tree.head)
}

func (tree *BinaryTree) delete(arg int, node *TreeNode) *TreeNode {
	switch {
	case arg < node.value:
		node.left = tree.delete(arg, node.left)
		return node
	case node.value < arg:
		node.right = tree.delete(arg, node.right)
		return node
	case node.value == arg:
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			node.right = node.right.lift(node)
		}
	}
	return nil
}

func (node *TreeNode) lift(nodeToDelete *TreeNode) *TreeNode {
	if node.left != nil {
		node.left = node.left.lift(nodeToDelete)
		return node
	} else {
		nodeToDelete.value = node.value
		return node.right
	}
}

func (tree *BinaryTree) TraverseAndPrint() {
	tree.head.TraverseAndPrint()
}

func (node *TreeNode) TraverseAndPrint() {
	if node == nil {
		return
	}
	node.left.TraverseAndPrint()
	fmt.Println(node.value)
	node.right.TraverseAndPrint()
}

func NewBinaryTree(arg int) *BinaryTree {
	return &BinaryTree{&TreeNode{left: nil, right: nil, value: arg}}
}

func (tree *BinaryTree) GreatestValue() int {
	return tree.head.greatestValue()	
}

func (node *TreeNode) greatestValue() int {
	if node.right != nil {
		return node.right.greatestValue()
	} else {
		return node.value
	}
}

func main() {
	tree := NewBinaryTree(50)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(10)
	tree.Insert(33)
	tree.Insert(56)
	tree.Insert(89)

	tree.Delete(50)

	tree.TraverseAndPrint()
	fmt.Println(tree.GreatestValue())

}

