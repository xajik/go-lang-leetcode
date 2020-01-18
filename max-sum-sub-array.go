package main

import "fmt"

var input = []int {10, 1, 3, 15, 30, 40, 4, 50, 2, 1, 6}
var K = 3
var L = 2

//SubArray with Sum of elements
type SubArray struct {
	sum int
	array []int
}

func getMaxSubArray(array []int, subArrayLenght int) []int {
	maxSum := 0
	minSumIndex := 0
	maxSumIndex := 0
	for i := 0; i < len(array) - subArrayLenght; i++ {
		sum := 0
		for j := i; j < i + subArrayLenght; j++ {
			sum += array[j]
		}
		if(maxSum < sum) {
			maxSum = sum
			minSumIndex = i;
			maxSumIndex = i + subArrayLenght;
		}
	}
	return array[minSumIndex:maxSumIndex]
}

func getMaxSubArrays(array []int, subArrayLenghts []int, startIndex int, output *[]SubArray) {
	if(startIndex == len(subArrayLenghts) || len(array) < subArrayLenghts[startIndex]) {
		return 
	}
	maxSum := 0
	minSumIndex := 0
	maxSumIndex := 0
	subArrayLenght := subArrayLenghts[startIndex]
	for i := 0; i < len(array) - subArrayLenght + 1; i++ {
		sum := 0
		for j := i; j < i + subArrayLenght; j++ {
			sum += array[j]
		}
		if(maxSum < sum) {
			maxSum = sum
			minSumIndex = i;
			maxSumIndex = i + subArrayLenght;
		}
	}
	if(maxSum != 0) {
		existingValue := (*output)[startIndex]
		if(existingValue.sum < maxSum) {
			newValue := SubArray{maxSum, array[minSumIndex:maxSumIndex]}
			(*output)[startIndex]  = newValue
		}
	}
	startIndex++
	getMaxSubArrays(array[0:minSumIndex], subArrayLenghts, startIndex, output)
	getMaxSubArrays(array[maxSumIndex:len(array)], subArrayLenghts, startIndex, output)
}

func testMaxSubArray() {
	fmt.Println()
	fmt.Println("Maximum sum of one subarrays of given size")
	fmt.Println("Input: ", input)
	subArray := getMaxSubArray(input, K)
	fmt.Printf("Outcome: lenght: %d; array: %v \n", K, subArray)
}

func testMaxSubArrays() {
	fmt.Println()
	fmt.Println("Maximum sum two non-overlapping subarrays of given size")
	subArrayLenght := []int{K, L}
	output := make([]SubArray, len(subArrayLenght))
	getMaxSubArrays(input, subArrayLenght, 0, &output)
	printSubArray(output)
}

func printSubArray(items []SubArray) {
	fmt.Println("Outcome:")
	for i, item := range items {
		fmt.Printf("[%d] sum: %d; items: %v; \n", (i + 1), item.sum, item.array)
	}
	fmt.Println()
}