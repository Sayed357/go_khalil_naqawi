package main

import "fmt"

func main() {

	// calculates and prints the sum of prime numbers for each provided slice
	// calculates the sum of prime numbers in each slice.

	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	var sum int = 0

	for _, num := range numbers {
		if checkPrime(num) {
			sum += num
		}
	}

	return sum
}

// determine if the number is prime.
func checkPrime(number int) bool {
	if number == 1 {
		return false // not a prime number
	}

	if number == 2 {
		return true // a prime number
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
