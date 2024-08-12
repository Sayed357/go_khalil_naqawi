package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {

	// your code here
	found := false
	for x := 1; x <= a; x++ {
		for y := 1; y <= a; y++ {
			for z := 1; z <= a; z++ {
				if (x+y+z == a) && (x*y*z == b) && (x*x+y*y+z*z == c) {
					fmt.Println(x, y, z)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		fmt.Println("Problem can not be solved")
	}

}

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

}
