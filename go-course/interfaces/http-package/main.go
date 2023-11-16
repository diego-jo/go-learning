package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.Get("https://google.com")
	fmt.Println(*r)
}
