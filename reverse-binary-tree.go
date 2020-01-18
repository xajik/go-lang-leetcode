package main

import (
	"fmt"
)

//https://leetcode.com/problems/invert-binary-tree/

//ReverseTreeNode 
type ReverseTreeNode struct {
	Val int
	Left *ReverseTreeNode
	Right *ReverseTreeNode
}

func createNode(key int) * ReverseTreeNode {
	return &ReverseTreeNode{key, nil, nil}
}

func insertTreeNode(root * ReverseTreeNode, key int) * ReverseTreeNode {
	if(root == nil) {
		return createNode(key)
	}
	if(key > root.Val) {
		root.Right = insertTreeNode(root.Right, key)
	} else {
		root.Left = insertTreeNode(root.Left, key)
	}
	return root
}

func invertTree(root *ReverseTreeNode) *ReverseTreeNode {
    if(root == nil) {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left) 
	invertTree(root.Right)
	return root
}

func printReversTree(root *ReverseTreeNode) {
	if(root == nil) {
		return 
	}
	fmt.Printf("%d,", root.Val)
	printReversTree(root.Left) 
	printReversTree(root.Right)
}

func testInvertTree() {
	fmt.Println()
	fmt.Println("Invert Tree Test")
	//[4,2,7,1,3,6,9]
	tree := createNode(4)
	insertTreeNode(tree, 2)
	insertTreeNode(tree, 7)
	insertTreeNode(tree, 1)
	insertTreeNode(tree, 3)
	insertTreeNode(tree, 6)
	insertTreeNode(tree, 9)

	fmt.Printf("\nBefore: [")
	printReversTree(tree)
	fmt.Printf("]\n")

	tree = invertTree(tree)

	fmt.Printf("\nAfter[")
	printReversTree(tree)
	fmt.Printf("]\n")
}