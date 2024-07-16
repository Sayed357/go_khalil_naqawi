package main

import (
	"fmt"
)

// check if a numebr is Prime
func primeNumber(number int) bool {

	// Numbers less than or equal to 1 are not prime
	// Numbers 2 and 3 are prime.
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}

	// If the number is divisible by 2 or 3, it is not prime
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true

}

func main() {

	// Declare "num" variable
	var num int  
	fmt.Print("Dail a number to check if it is a Prime: ")

	// check if the input was read accuratly.
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if primeNumber(num) {
		fmt.Printf("%d Prime number.", num)
	} else {
		fmt.Printf("%d Not a Prime number.", num)
	}
}
