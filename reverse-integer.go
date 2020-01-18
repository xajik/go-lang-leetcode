package main

//https://leetcode.com/problems/reverse-integer/

import (
	"fmt"
)

func reverse(x int) int {
	res := 0
	sign := 1
	if(x < 0) {
		sign = -1
	}

	x = Abs(x)
	for {
		tmp := x % 10
		res += tmp
		x = x/10
		fmt.Printf("tmp : %d; res: %d; x: %d \n", tmp, res, x)
		if(x <= 0) {
			break 
		}
		res *= 10
		//sign Max Int 32
		if(res > 2147483647) {
			return 0
		}
	}
	return res * sign
}

//Abs for
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func testRevertInt() {
	fmt.Println()
	fmt.Println("Revert Integer test")

	number := 120
	fmt.Printf("Number : %d; reversed: %d; \n", number, reverse(number))
	number = 123
	fmt.Printf("Number : %d; reversed: %d; \n", number, reverse(number))
	number = -123
	fmt.Printf("Number : %d; reversed: %d; \n", number, reverse(number))
	number = 1534236469
	fmt.Printf("Number : %d; reversed: %d; \n", number, reverse(number))
}