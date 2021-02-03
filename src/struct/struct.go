package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "Person One", age: 42}
	p2 := Person{name: "Person Two", age: 22}

	fmt.Println(p1.age - p2.age)

	var weird []map[string]int

	a := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	weird = append(weird, a)

	fmt.Println(weird[0]["one"])
}
