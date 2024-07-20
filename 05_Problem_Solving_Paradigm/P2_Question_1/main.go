package main

import (
	"fmt"
	"strings"
)

// Roman numeral symbols and its corresponding values
var (
	symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func intToRoman(num int) string {
	// Init result to store the Roman numeral
	var result strings.Builder

	for i := 0; i < len(values); i++ {
		// Repeat the current symbol while its value can be subtracted from num
		for num >= values[i] {
			num -= values[i]
			result.WriteString(symbols[i])
		}
	}

	return result.String()
}

func main() {
	fmt.Println(intToRoman(4))

	fmt.Println(intToRoman(9))

	fmt.Println(intToRoman(23))

	fmt.Println(intToRoman(2021))

	fmt.Println(intToRoman(1646))
}
