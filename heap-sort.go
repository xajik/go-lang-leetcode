package main

import (
    "fmt"
)

//MaxUint Value
const MaxUint = ^uint(0) 
//MinUint Value
const MinUint = 0 
//MaxInt Value
const MaxInt = int(MaxUint >> 1) 
//MinInt Value
const MinInt = -MaxInt - 1

//Heap structure
type Heap struct {
	heapArray []int
	size int
	maxSize int
}

func createHeap(maxSize int) *Heap {
	p := new(Heap)
	p.size = 0 
	p.maxSize = maxSize
	p.heapArray = make([]int, maxSize + 1)
	p.heapArray[0] = MaxInt
	return p
}

func createTenHeapPointer() * Heap {
	return &Heap{make([]int, 10), 0, 10}
}

func createTenHeapObject() Heap {
	return Heap{make([]int, 10), 0, 10}
}

func (Heap) parent(position int) int {
	return position/2
}

func (Heap) leftChild(position int) int {
	return position * 2
}

func (Heap) rightChild(position int) int {
	return (position * 2) + 1
}

func (heap Heap) isLeaf(position int) bool {
	return position >= (heap.size / 2) && position <= heap.size
} 

func (heap Heap) swap(p1 int, p2 int) {
	tmp := heap.heapArray[p1]
	heap.heapArray[p1] = heap.heapArray[p2]
	heap.heapArray[p2] = tmp
} 

// A recursive function to max heapify the given 
// subtree. This function assumes that the left and 
// right subtrees are already heapified, we only need 
// to fix the root. 
func (heap Heap) maxHeapify(position int) {
	if(heap.isLeaf(position)) {
		return 
	}
	value := heap.heapArray[position]
	leftChildPosition := heap.leftChild(position)
	rightChildPosition := heap.rightChild(position)
	leftChild := heap.heapArray[leftChildPosition]
	rightChild := heap.heapArray[rightChildPosition]
	if(value < leftChild || value < rightChild) {
		if(leftChild > rightChild) {
			heap.swap(position, leftChildPosition)
			heap.maxHeapify(leftChildPosition)
		} else {
			heap.swap(position, rightChildPosition)
			heap.maxHeapify(rightChildPosition)
		}
	}
}

func (heap *Heap) insert(value int) {
	heap.size ++
	heap.heapArray[heap.size] = value
	currectSize := heap.size
	for (heap.heapArray[currectSize] > heap.heapArray[heap.parent(currectSize)]) {
		heap.swap(currectSize, heap.parent(currectSize))
		currectSize = heap.parent(currectSize)
	}
}

func (heap Heap) printHeap() {
	for i := 1; i < heap.size/2; i++ {
		fmt.Printf("\nParent: %d", heap.heapArray[i])
		fmt.Printf(";\t Left: %d", heap.heapArray[heap.leftChild(i)])
		fmt.Printf(";\t Right: %d", heap.heapArray[heap.rightChild(i)])
	}
	fmt.Println("")
}

func (heap *Heap) extractmax() int {
	if heap.size <= 1 {
		return -1
	}
 	max := heap.heapArray[1]
	heap.size --
	heap.heapArray[1] = heap.heapArray[heap.size]
	heap.maxHeapify(1)
	return max
}

func testHeapSort() {
	fmt.Println("***")
	fmt.Println("Heap Data Structure")
	heap := createHeap(15)
	fmt.Printf("\nHeap: size %d ; max size %d ; array lenght: %d ", heap.size, heap.maxSize,  len(heap.heapArray))
	heap.insert(5) 
	heap.insert(3)
	heap.insert(17) 
	heap.insert(10) 
	heap.insert(84) 
	heap.insert(19) 
	heap.insert(6)
	heap.insert(22) 
	heap.insert(9)
	heap.printHeap()
	fmt.Println("Max heap value: ", heap.extractmax())	
}
