package main

import "fmt"

func main() {

	printRotatedPyramid()
	fmt.Println()
	printPyramid()
	a, b := returnIntAndFloat(1, 1.1)
	fmt.Println(a, b)

}

func printRotatedPyramid() {
	for i := 0; i < 5; i++ {
		if i < 3 {
			for j := 0; j < i+1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		} else {
			for j := 0; j < 5-i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func printPyramid() {
	leftBorder := 2
	rightBorder := 4
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if j > leftBorder && j < rightBorder {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		if i < 3 {
			leftBorder--
			rightBorder++
		} else {
			leftBorder++
			rightBorder--
		}
		fmt.Println()

	}
}

func returnIntAndFloat(num1 int, num2 float64) (int, float64) {
	return num1, num2
}
