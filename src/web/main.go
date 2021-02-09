package main

import (
	"Go/src/myapp"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is up listening on PORT: 8080...")
	http.ListenAndServe(":3001", myapp.NewHttpHandler())
}
