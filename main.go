package main

import "fmt"

/**
* Binary Search Trees
* Tutorial from https://levelup.gitconnected.com/binary-search-trees-in-go-58f9126eb36b
* The left child node is always smaller than the right child node.
*
* Node		: Each node has a left & right pointer to its children
* Leaf Node	: A node with no children
* Root Node	: topmost node
* Edge		: Links two nodes
**/

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return fmt.Errorf("tree is nil")
	}

	if t.value == value {
		return fmt.Errorf("this node value already exists")
	}

	// if the new value is smaller than the current value, insert it as a left child node.
	// if however, the node already has a left child node, save the value as a left child node
	// recursively using this function.
	if t.value > value {
		if t.left == nil {
			t.left = &TreeNode{value: value}
			return nil
		}
		return t.left.Insert(value)
	}

	// if the new value is is greater than the current node value, insert it on the right.
	// if the right child node already exists, do an insert on the right child node
	// recursively using this function.
	if t.value < value {
		if t.right == nil {
			t.right = &TreeNode{value: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindMin() int {
	// returns the leftmost leaf node
	// because a bst is already sorted, the left most leaf node is the minimum value
	if t.left == nil {
		return t.value
	}
	return t.left.FindMin()
}

func (t *TreeNode) FindMax() int {
	// returns the rightmost leaf node
	// because a bst is already sorted, the right most leaf node is the maximum value
	if t.right == nil {
		return t.value
	}
	return t.right.FindMax()
}

func (t *TreeNode) PrintInOrder() {
	// in order traversal will always be in ascending order.
	// the left child is visited first, then the paren and then the right child.
	if t == nil {
		return
	}
	t.left.PrintInOrder()
	fmt.Println(t.value)
	t.right.PrintInOrder()
}

func main() {
	var tree = &TreeNode{value: 8}
	tree.Insert(34)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(6)

	fmt.Printf("min : %d\n", tree.FindMin())
	fmt.Printf("max : %d\n", tree.FindMax())
	tree.PrintInOrder()
}
