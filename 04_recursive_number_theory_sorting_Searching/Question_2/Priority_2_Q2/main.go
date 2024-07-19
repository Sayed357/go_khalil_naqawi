package main

import (
	"fmt"
)

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	if n <= 1 {
		fmt.Println("Not a Prime Factor")
		return
	}

	factors := []int{}
	// divide by 2 until "n" is an odd number
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// divide by odd numbers from 3 above
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// if "n" is ">" than 2, it is prime number
	if n > 2 {
		factors = append(factors, n)
	}

	// here we print the factors:
	for i, factor := range factors {
		if i == len(factors)-1 {
			fmt.Printf("%d\n", factor)
		} else {
			fmt.Printf("%d, ", factor)
		}
	}
}
