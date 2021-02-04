package main

import "fmt"

type eater struct {
	hungryLevel int
	eaten       []string
}

func (eater *eater) eat(ate string) {
	eater.hungryLevel++
	eater.eaten = append(eater.eaten, ate)
}

func main() {
	eater1 := eater{hungryLevel: 40}
	fmt.Println(eater1.hungryLevel)
	eater1.eat("hamburger")
	eater1.eat("pizza")
	eater1.eat("pasta")
	fmt.Println(eater1.hungryLevel, eater1.eaten)
	eater1.eat("steak")
	fmt.Println(eater1.hungryLevel, eater1.eaten)
}
