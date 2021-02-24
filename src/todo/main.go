package main

import (
	"fmt"
	"net/http"

	"github.com/bombamong/go_study/src/todo/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	fmt.Println("Listening on localhost:3000")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
