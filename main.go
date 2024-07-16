package main

import (
	"fmt"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	// Calculate the number of positions to spin
	numbers := len(sentence)
	if numbers == 0 {
		return sentence
	}

	//  rotate by 1 position
	positions := 1

	// Spin the string by moving characters to the right
	return sentence[numbers-positions:] + sentence[:numbers-positions]
}
