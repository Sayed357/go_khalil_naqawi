package main

import (
	"fmt"
)

// convert an integer  to its binary string representation
func intToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var result string

	// convert remainder 0/1 and prepend to the result string
	for n > 0 {
		remainder := n % 2
		result = fmt.Sprint(remainder) + result
		n = n / 2
	}
	return result
}

// generate a list of binary representations from "0" to "n"
func generateBinaryList(n int) []string {
	var result []string
	for i := 0; i <= n; i++ {
		result = append(result, intToBinary(i))
	}
	return result
}

func main() {
	n := 2
	binaryList := generateBinaryList(n)
	fmt.Println(binaryList)

	n = 3
	binaryList = generateBinaryList(n)
	fmt.Println(binaryList)
}
