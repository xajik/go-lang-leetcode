package main

import (
	"math/rand"
	"fmt"
)

func findMedianSortedArraysSimple(array1 []int, array2 []int) {
	array := append(array1, array2...)
	fmt.Printf("Concatenated array: %v \n", array)
	array = quickSort(array)
	fmt.Printf("Sorted concatenated array: %v \n", array)
	medianIndex := 0
	median := 0.0
	if (len(array) % 2 == 0) {
		medianIndex = len(array) / 2
		median = (float64(array[medianIndex - 1]) + float64(array[medianIndex])) / 2
	} else {
		medianIndex = len(array) / 2
		median = float64(array[medianIndex])
	}
	fmt.Printf("Median [%d] %.2f \n", medianIndex, median)
}

func quickSort(a []int) []int {
	if len(a) < 2 { 
		return a 
	}
	left:= 0 
	right := len(a) - 1
	// Pick a pivot
	pivotIndex := rand.Int() % len(a)
	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	// Pile elements smaller than the pivot on the left
	for i := range a {
	  if a[i] < a[right] {
		//swap
		a[i], a[left] = a[left], a[i]
		left++
	  }
	}
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
	// Go down the rabbit hole
	quickSort(a[:left])
	quickSort(a[left + 1:])
	return a
  }

func testMedianSortedArraysSimple() {
	fmt.Println()
	fmt.Printf("Median Sorted Arrays Simple")
	findMedianSortedArraysSimple([]int{1,3}, []int{2})
	fmt.Println()
	findMedianSortedArraysSimple([]int{3,4}, []int{1,2})
}

//Complexity: O(log(min(m,n))).
//Task from: https://leetcode.com/problems/median-of-two-sorted-arrays/
//Solution from: https://medium.com/@hazemu/finding-the-median-of-2-sorted-arrays-in-logarithmic-time-1d3f2ecbeb46

func findMedianSortedArraysFast(a []int, b []int)  float32 {
	if(a == nil || b == nil) {
		fmt.Printf("Invalid arrays. A: %v; B: %v \n", a, b)
		return 0.0
	}
	fmt.Printf("Input arrays. A: %v; B: %v \n", a, b)

	lenghtA := len(a)
	lenghtB := len(b)
	leftHalfLenght := (lenghtA + lenghtB + 1) / 2
	fmt.Printf("Left half lenght: %d \n", leftHalfLenght)

	//Make sure A is shorted
	if(lenghtA > lenghtB) {
		//Swap lenght and array
		a, b = b, a
		lenghtA, lenghtB = lenghtB, lenghtA
	}
	fmt.Printf("Lenght of arrays after swap: A[%d]; B[%d] \n",lenghtA, lenghtB)

	aMinCount := 0
	aMaxCount := lenghtA

	for aMinCount <= aMaxCount {
		aCount := aMinCount + (aMaxCount - aMinCount) / 2
		bCount := leftHalfLenght - aCount
		
	    if(aCount > 0 && a[aCount - 1] > b[bCount]) {
			aMaxCount = aCount - 1
		} else if(aCount < lenghtA && b[bCount - 1] > a[aCount]) {
			aMinCount = aCount + 1
		} else {
			leftSideEnd := func() int {
				if(aCount == 0) {
					 return b[bCount - 1]
				} else if(bCount == 0) {
					return a[aCount - 1]
				} 
				if(a[aCount - 1] > b[bCount - 1]) {
					return a[aCount - 1]
				}
				//else 
				return b[bCount - 1]
				}()
			
			
			//If number is odd - we else found element found 
			if(isOdd(lenghtA + lenghtB)) {
				fmt.Printf("Is Odd result: %d \n", leftSideEnd)
				return float32(leftSideEnd)
			}

			rightSideStart := func() int {
					if(aCount == lenghtA) {
						return b[bCount]
					} else if(bCount == lenghtB) {
						return a[aCount]
					}
					if(a[aCount] < b[bCount]) {
						return a[aCount]
					}
					return b[bCount]
				}()

			//for NOT odd number we need to find a middle manually
			result := (float32(leftSideEnd) + float32(rightSideStart)) / 2
			fmt.Printf("Is not Odd result: %.2f \n", result)
			return result
		}
	}

	panic("This code should not be reached. Method should finish in 'else' ")
  }

  //The least significant bit of any odd number is 1
func isOdd(number int) bool {
	  return number & 1 == 1
  }

func testMedianSortedArraysFast() {
	fmt.Println()
	fmt.Printf("Median Sorted Arrays Fast \n")
	findMedianSortedArraysFast([]int{1,3}, []int{2})
	fmt.Println()
	findMedianSortedArraysFast([]int{3,4}, []int{1,2})
}