package main

import (
	"fmt"
	"slices"
)

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) // [1,2,5,4,3,7,8]

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) // [1, 2, 5, 4, 7, 9, 10]

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) // [1, 4, 5, 7]
}

func merge(data [][]int) []int {
	arr1 := data[0]
	arr2 := data[1]
	result := []int{}    // to store the final result.
	tmp := map[int]int{} // map to store the frequency of each integer.

	var combined []int

	// Elements from arr1 and arr2 are appended to combined
	combined = append(combined, arr1...)
	combined = append(combined, arr2...)

	// calculate the frequencies
	for _, val := range combined {
		_, ok := tmp[val]
		if ok {
			tmp[val]++
		} else {
			tmp[val] = 1
		}
	}

	// inside the map -> key must be unique

	for key := range tmp {
		result = append(result, key)
	}

	// Sort the result slice using slices.Sort
	slices.Sort(result)

	return result
}
