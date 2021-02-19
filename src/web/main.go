package main

import (
	"fmt"
	"net/http"

	"github.com/bombamong/go_study/src/myapp"
)

//PORT of server
const PORT string = "3000"

func main() {

	fmt.Printf("Server is up listening on PORT: %s...\n", PORT)
	http.ListenAndServe(":"+PORT, myapp.NewHttpHandler())
}
