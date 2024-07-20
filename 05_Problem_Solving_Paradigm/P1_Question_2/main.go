package main

import (
	"fmt"
)

// create pascal triangle with the given number od row
func generatePascalTriangle(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	// init a 2D slice with the given number of row

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)

		// set the first and last elements in between the 1st and last elements.
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {

	// define the number of row
	numRows := 5

	pascalTriangle := generatePascalTriangle(numRows)
	for _, row := range pascalTriangle {
		fmt.Println(row)
	}
}
