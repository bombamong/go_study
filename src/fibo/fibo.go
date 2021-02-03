package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Printf("Fibonacci for %d is %d \n", i, fibo(i))
	}
}

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-2) + fibo(x-1)
}
