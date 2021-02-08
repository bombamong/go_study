package main

import "fmt"

func main() {
	var c chan int
	c = make(chan int)

	go pop(c)

	c <- 10
	c <- 20
	c <- 30

	fmt.Println("end of program")
}

func pop(c chan int) {
	for {
		fmt.Println("POP")
		v := <-c
		fmt.Println(v)
	}
}
