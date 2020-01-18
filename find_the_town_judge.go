package main

import (
    "fmt"
)

//Task from: https://leetcode.com/problems/find-the-town-judge/

func findJudge(N int, trust [][]int) int {
	if(len(trust) == 0) {
		return N
	}
	trustIndex := -1
	maxValue := -1
	trustReceived := make([]int, N)
	trustGave := make([]int, N)
	for i := 0; i < len(trust); i++ {
		trusted := trust[i][1] - 1
		trustReceived[trusted] ++
		trustGave[trust[i][0] - 1]++
		if(trustReceived[trusted] > maxValue) {
			maxValue = trustReceived[trusted]
			trustIndex = trusted
		}
	}
	if(trustIndex == -1 || maxValue != N - 1) {
		return -1
	}
	if(trustGave[trustIndex] > 0) {
		return -1
	}
	return trustIndex + 1
}

func testFindJudge() {
	fmt.Println()
	fmt.Println("Find Judge Task")
	N := 4
	trust := [][]int{{1,3},{1,4},{2,3},{2,4},{4,3}}
	result := findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n", N, trust, result, (result == 3))

	N = 3 
	trust = [][]int{{1,2},{2,3}}
	result = findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n" , N, trust, result, (result == -1))

	N = 2
	trust = [][]int{{1,2}}
	result = findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n", N, trust, result, (result == 2))

	N = 2
	trust = [][]int{{1,2}}
	result = findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n", N, trust, result, (result == 2))

	N = 3
	trust = [][]int{{1,3},{2,3},{3,1}}
	result = findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n", N, trust, result, (result == -1))

	N = 1
	trust = [][]int{}
	result = findJudge(N, trust)
	fmt.Printf("Find Judge: N: %d ; input: %v ; result: %d [%t] \n", N, trust, result, (result == 1))

}