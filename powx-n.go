package main

//https://leetcode.com/problems/powx-n/

import (
    "fmt"
    //"math"
)

func myPow(x float64, n int) float64 {
    if(n == 0) {
        return 1
    }
    if(n == 1) {
        return x
    }
    if(x == 1) {
        return 1
    }
    if(x == -1) {
        if(int(n) % 2 == 0) {
            return 1
        }
        return -1
    }

    if(n < 0) {
        x = 1/x
    }

    if(n < 0 ) { 
        n = -n 
    }

    res := x
    for i := 1; i < n; i++ {
        res *= x
    }

    return res
}

func testPow() {
    fmt.Println("***")
    fmt.Println("Test My Pow")
    
    number := 2.00000
    pow := 2
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))

    number = 2.00000
    pow = 10
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))
    
    number = 2.10000
    pow = 3
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))

    number = 2.00000
    pow = -2
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))

    number = 1.00000
    pow = -2147483648
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))

    number = -1.00000
    pow = 2147483648
    fmt.Printf("Number %f to pow %d: %f; \n", number, pow, myPow(number, pow))


}