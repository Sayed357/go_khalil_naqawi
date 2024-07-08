// +Question 1

// package main

//import (
// "fmt"
//"math"
// )

// func main() {
// const Phi = 3.14

// var outerRadius, innerRadius, height float64

// fmt.Print("input the outer radius of the tube: ")
// fmt.Scan(&outerRadius)

// fmt.Print("input the inner radius of the tube: ")
// fmt.Scan(&innerRadius)

// fmt.Print("input the height of the tube: ")
// fmt.Scan(&height)

// volume := Phi * (math.Pow(outerRadius, 2) - math.Pow(innerRadius, 2)) * height

// fmt.Printf("The volume of the tube is: %.2f\n", volume)
// }

// Question 2

// package main

// import "fmt"

// func main() {
// var score int

// fmt.Print("Dail the score: ")
// fmt.Scan(&score)

// if score < 0 || score > 100 {
// fmt.Println("Invalid score")
//} else if score >= 85 && score <= 100 {
// fmt.Println("Grade: A")
// } else if score >= 70 && score <= 84 {
// fmt.Println("Grade: B")
// } else if score >= 55 && score <= 69 {
// fmt.Println("Grade: C")
// } else if score >= 40 && score <= 50 {
// fmt.Println("Grade: D")
// } else {
// fmt.Println("Grade: E")
// }
// }

// Question 3

// package main

// import (
// "fmt"
// "math"
// )

// func main() {
// for i := 1; i <= 100; i++ {
// if i%4 == 0 && i%7 == 0 {
// fmt.Println("buzz")
// } else if i%4 == 0 {
// fmt.Println(int(math.Pow(float64(i), 2)))
// } else if i%7 == 0 {
// fmt.Println(int(math.Pow(float64(i), 3)))
// } else {
// fmt.Println(i)
//}
//}
//}

// Question 3 Priority 2

package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func isAnagram(word1, word2 string) bool {

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	sortedWord1 := sortString(word1)
	sortedWord2 := sortString(word2)

	return sortedWord1 == sortedWord2
}

func main() {
	var word1, word2 string

	word1 = "cafe"
	word2 = "face"
	if isAnagram(word1, word2) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not Anagram")
	}

	word1 = "donut"
	word2 = "peanut"
	if isAnagram(word1, word2) {
		fmt.Println("Not Anagram")
	} else {
		fmt.Println("Not Anagram")
	}
}
