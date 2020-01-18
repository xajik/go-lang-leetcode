package main

import "fmt"

//TreeNode node
type TreeNode struct {
	value int
	rightNode *TreeNode
	leftNode *TreeNode
}

func newNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func insert(root *TreeNode, value int) *TreeNode {
	if(root == nil) {
		return newNode(value)
	}
	if(value > root.value) {
		root.rightNode = insert(root.rightNode, value)
	} else if(value < root.value) {
		root.leftNode = insert(root.leftNode, value)
	}
	return root
}

func printTree(root *TreeNode) {
	if(root != nil) {
		fmt.Println("Value: ", root.value)
	}
	if(root.leftNode != nil) {
		printTree(root.leftNode)
	}
	if(root.rightNode != nil) {
		printTree(root.rightNode)
	}
}

func treeLenght(lenght *int, root *TreeNode) int {
	if(root != nil) {
		*lenght ++
	}
	if(root.leftNode != nil) {
		treeLenght(lenght, root.leftNode)
	}
	if(root.rightNode != nil) {
		treeLenght(lenght, root.rightNode)
	}
	return *lenght
}

func testTree() {
	fmt.Println("***")
	fmt.Println("Tree Data Structure")
	root := newNode(10)
	values := []int{4,6,1,2,9,0,15}
	for _, v := range values {
		insert(root, v)
	}
	lenght := 0
	fmt.Printf("\nTree lenght: %d \n", treeLenght(&lenght, root))
	fmt.Printf("\nTree is balanced: %t \n", isBalanced(root))
	printTree(root)
}

func isBalanced(root *TreeNode) bool {
	leftBranch := 0
	rightBranch := 0
	return treeLenght(&leftBranch, root.leftNode) == treeLenght(&rightBranch, root.rightNode)
} 



