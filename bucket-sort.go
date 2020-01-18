package main

import "fmt"

//SlotSize is # 10 means each slot's size is 0.1 
const SlotSize = 10

//Array wrapper for the 2d array
type Array struct {
	array[] float32
}  

func (a *Array) insertionSort() {
	array := a.array
	for i := 1; i < len(array); i++ {
		value := array[i]
		j := i - 1

		for j >= 0 && array[j] > value {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = value
	}
}

func bucketSort(array []float32) []float32 {
	buckets := make([]Array, SlotSize)

	for i := 0; i <  len(array); i++ {
		var subIndex int = int(array[i] * SlotSize)
		bucket := &buckets[subIndex]
		bucket.array = append(bucket.array, array[i])
	}

	for i := range buckets {
		buckets[i].insertionSort()
	}

	var sortedBuckets []float32
	for i := range array {
		sortedBucket := buckets[i].array
		sortedBuckets = append(sortedBuckets, sortedBucket...)
	}
	return sortedBuckets
}

func print(buckets []float32) {
	for i := range buckets {
		fmt.Println(buckets[i])
	}
}

func testBucketSort() {
	fmt.Println("***")
	fmt.Println("Bucket sort")
	fmt.Println()
	array := []float32{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	buckets := bucketSort(array)
	print(buckets)
}

