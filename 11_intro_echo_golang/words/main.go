package main

import (
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	words []string
	mu    sync.Mutex
)

type WordRequest struct {
	Word string `json:"word"`
}

type WordResponse struct {
	Word        string `json:"word"`
	Length      int    `json:"length"`
	NumOfVocals int    `json:"num_of_vocals"`
	Palindrome  bool   `json:"palindrome"`
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	e := echo.New()

	e.POST("/words", func(c echo.Context) error {
		var wordReq WordRequest

		if err := c.Bind(&wordReq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "invalid request",
			})
		}

		word := wordReq.Word
		length := len(word)
		numOfVocals := countVowels(word)
		palindrome := isPalindrome(word)

		// insert new word with mutex lock
		mu.Lock()
		words = append(words, word)
		mu.Unlock()

		response := WordResponse{
			Word:        word,
			Length:      length,
			NumOfVocals: numOfVocals,
			Palindrome:  palindrome,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "word added",
			"data":    response,
		})
	})

	e.GET("/words", func(c echo.Context) error {
		mu.Lock()
		defer mu.Unlock()
		return c.JSON(http.StatusOK, echo.Map{
			"message": "all data",
			"data":    words,
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
