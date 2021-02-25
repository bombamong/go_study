package main

import (
	"fmt"
	"net/http"

	"github.com/bombamong/go_study/src/todo/app"
)

func main() {
	m := app.MakeNewHandler("./test.db")
	defer m.Close()

	fmt.Println("Listening on localhost:3000")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
