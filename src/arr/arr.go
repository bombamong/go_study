package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := arr
	arr2[0] = 2
	fmt.Println(arr, arr2)

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice
	slice2[0] = 2
	fmt.Println(slice, slice2)
}

func learnOne() {
	s := "Hello 월드"
	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
	fmt.Println()

	for _, c := range s {
		fmt.Print(string(c), ", ")
	}
	fmt.Println()

	fmt.Println(rand.Intn(1))
}
