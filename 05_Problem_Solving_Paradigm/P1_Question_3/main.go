package main

import "fmt"

func fibonacci(n int) int {
	// your code here

	// base cases
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	// Init fibonacci sequence array with the first two numbers
	table := []int{0, 1}

	// Build fibonacci sequence up to the n number
	for i := 2; i <= n; i++ {
		table = append(table, table[i-1]+table[i-2])
	}
	return table[n]
}

func main() {
	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144
}
