package main

import (
	"fmt"
	"example.com/nums"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println(addTwoNumbers(1, 2))
	fmt.Println(nums.AddThreeNumbers(1,2,3))
}

func addTwoNumbers (num1 int, num2 int) int {
	return num1 + num2
}