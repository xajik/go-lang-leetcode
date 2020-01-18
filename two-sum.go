package main

//https://leetcode.com/problems/two-sum/

import "fmt"

func twoSumDoubleLoop(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if(i == j) {
				continue
			}
			if(nums[i] + nums[j] == target) {
				return []int{i,j}
			}
		}
	}
	
	return []int{-1,-1}
}

func twoSumSimpleMap(nums []int, target int) []int {
	diffMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if val, ok := diffMap[diff]; ok {
			return []int{val, i}
		}
		diffMap[num] = i
	}
	return []int{-1,-1}
}

func testTwoSum() {
	fmt.Println("***")
	fmt.Println("Two Sum")

	target := 9
	nums := []int{2, 7, 11, 15}
	result := twoSumDoubleLoop(nums, target)
	fmt.Printf("Loop result. Input: %v; target: %d; result: %v \n", nums, target, result)

	target = 9
	nums = []int{2, 7, 11, 15}
	result = twoSumSimpleMap(nums, target)
	fmt.Printf("Map result. Input: %v; target: %d; result: %v \n", nums, target, result)

}


