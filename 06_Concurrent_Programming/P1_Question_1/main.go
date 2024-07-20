package main

import (
	"fmt"
	"time"
)

// print each step of the word being reversed with a 3s interval
func reverseWord(word string, done chan bool) {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {

		reversed += string(word[i])
		fmt.Println(reversed)
		time.Sleep(3 * time.Second)
	}

	// sign that the goroutine is done
	done <- true
}

func main() {
	word := "Phone"
	done := make(chan bool)

	go reverseWord(word, done)

	// wait for the goroutine to signal complete
	<-done
}
