package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

/*
Answer that the user needs to guess
*/
type Answer struct {
	answer1 int
	answer2 int
	answer3 int
}

/*
GenerateNewNumbers creates new random integers
and assigns to Answer
*/
func (a *Answer) GenerateNewNumbers() {
	a.answer1 = rand.Intn(9)
	a.answer2 = rand.Intn(9)
	a.answer3 = rand.Intn(9)
}

/*
ReceiveAnswer from user input.
*/
func (a *Answer) ReceiveAnswer() {
	var answers [3]int

	for i := 0; i < len(answers); i++ {
		fmt.Printf("Please enter guess #%d\n", i+1)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		answers[i], _ = strconv.Atoi(line)
	}
	a.answer1, a.answer2, a.answer3 = answers[0], answers[1], answers[2]
}

func checkAnswers(ca, ha Answer) bool {
	caAnswers := [3]int{ca.answer1, ca.answer2, ca.answer3}
	haAnswers := [3]int{ha.answer1, ha.answer2, ha.answer3}

	fmt.Println("\n==================================\n")
	fmt.Printf("%d %d %d are your guesses\n", ha.answer1, ha.answer2, ha.answer3)
	fmt.Println("\n==================================\n")

	if ca.answer1 == ha.answer1 && ca.answer2 == ha.answer2 && ca.answer3 == ha.answer3 {
		fmt.Println("Congratulations, you won!")
		return true
	} else {
		counter := 0
		for i := 0; i < len(caAnswers); i++ {
			if caAnswers[i] == haAnswers[i] {
				counter++
			}
		}
		fmt.Printf("You got %d / 3 answers correct, try again!\n", counter)
		return false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	computerAnswer := Answer{}
	computerAnswer.GenerateNewNumbers()

	humanAnswer := Answer{}
	tries := 1
	for {
		humanAnswer.ReceiveAnswer()
		result := checkAnswers(computerAnswer, humanAnswer)
		if result == true {
			fmt.Printf("You have attempted %d times.", tries)
			break
		}
		tries++
	}
}
