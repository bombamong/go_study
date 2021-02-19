package main

import (
	"fmt"
	"net/http"

	"github.com/bombamong/go_study/src/restapi/myapp"
)

func main() {
	fmt.Println("running server on localhost:3000")
	http.ListenAndServe(":3000", myapp.NewHandler())
}
