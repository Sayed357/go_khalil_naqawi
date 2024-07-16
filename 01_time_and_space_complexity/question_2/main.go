package main

import (
	"fmt"
)

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}

func pow(x, n int) int {

	// base case n=0: if the exponent n is 0, the result is always 1.
	if n == 0 {
		return 1
	}

	// handing negative exponent
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// odd and even "n"
	if n%2 == 0 {
		half := pow(x, n/2)
		return half * half
	} else {
		return x * pow(x, n-1)
	}
}
