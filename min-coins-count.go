package main

import "fmt"

var coins = []int{1,2,5,10,25,50}

func countCoins(price int) int {
	leftPrice := price
	count := 0
	for i := len(coins)-1; i >= 0; i-- {
		if(leftPrice == 0) {
			break
		}
		coint := coins[i]
		if((leftPrice % coint) >= 0) {
			leftPrice = leftPrice % coint
			count ++
		}
	 }
	return count
}

func testPrice(price int) {
	count := countCoins(price)
	fmt.Printf("For %d you need %d coint \n", price, count)
}

func testCoins() {
	fmt.Println("***")
	fmt.Println("Coint count test")
	testPrice(99)
	testPrice(75)
	testPrice(77)
}