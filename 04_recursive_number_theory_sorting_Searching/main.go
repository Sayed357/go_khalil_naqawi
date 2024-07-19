package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	//TODO: your code here
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibX(n-1) + fibX(n-2) + 1
	}
}
