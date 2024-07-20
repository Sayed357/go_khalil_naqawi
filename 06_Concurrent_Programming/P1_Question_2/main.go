package main

import (
	"fmt"
)

// check if a number is Prime

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// goroutine to generate Prime numbers from 1-100
func generatePrimes(primes chan<- int) {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			primes <- i
		}
	}
	close(primes)
}

// calculate the sqaure of each Prime number
func calculateSquares(primes <-chan int, squares chan<- int) {
	for prime := range primes {
		squares <- prime * prime
	}
	close(squares)
}

func main() {
	primes := make(chan int)
	squares := make(chan int)

	go generatePrimes(primes)
	go calculateSquares(primes, squares)

	// Print the squares of prime numbers

	for square := range squares {
		fmt.Println(square)
	}

}
