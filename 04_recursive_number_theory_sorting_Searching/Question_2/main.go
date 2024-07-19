package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
	// output:
	// best student in math: aziz with 100
	// best student in science: jamie with 88
	// best student in english: jamie with 90
	// best student in average: aziz with 84

}

func getInfo(students []Student) {
	// TODO: your code here

	var bestMath, bestScience, bestEnglish, bestAverage Student
	var highestMath, highestScience, highestEnglish, highestAverage int

	for _, student := range students {
		// Check for the highest Math score
		if student.MathScore > highestMath {
			highestMath = student.MathScore
			bestMath = student
		}

		// Check for the highest Science score
		if student.ScienceScore > highestScience {
			highestScience = student.ScienceScore
			bestScience = student
		}

		// Check for the highest English score
		if student.EnglishScore > highestEnglish {
			highestEnglish = student.EnglishScore
			bestEnglish = student
		}

		// Calculate the average score
		averageScore := (student.MathScore + student.ScienceScore + student.EnglishScore) / 3

		// Check for the highest average score
		if averageScore > highestAverage {
			highestAverage = averageScore
			bestAverage = student
		}
	}

	fmt.Println("Best student in Math ", bestMath.Name, "with", bestMath.MathScore)
	fmt.Println("Best student in Science ", bestScience.Name, "with", bestScience.ScienceScore)
	fmt.Println("Best student in ENglish ", bestEnglish.Name, "with", bestEnglish.EnglishScore)
	fmt.Println("Highest average student ", bestAverage.Name, "with", highestAverage)

}
