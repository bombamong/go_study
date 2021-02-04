package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) printName() {
	fmt.Printf("My name is %s\n", p.name)
}

func main() {
	p1 := Person{name: "Person One", age: 42}
	p2 := Person{name: "Person Two", age: 22}

	fmt.Println(p1.age - p2.age)
	p1.printName()
	p2.printName()

	var weird []map[string]int
	a := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	weird = append(weird, a)
	fmt.Println(weird[0]["one"])

	animal1 := Animal{color: "brown"}
	animal1.ChangeColor("blue")
	fmt.Println(animal1.color)

	animal1.ReallyChangeColor("blue")
	fmt.Println(animal1.color)

	s := "안녕이건한글"
	sRune := []rune(s)
	sRune[2] = rune('삼')
	fmt.Println(string(sRune))
}
