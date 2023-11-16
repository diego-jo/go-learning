package main

import (
	"fmt"
	"io"
	"os"
)

type abobora struct{}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// a := &abobora{}

	io.Copy(os.Stdout, f)
}

// func (abobora) Write(p []byte) (n int, err error) {
// 	fmt.Println(string(p))
// 	return 1, nil
// }
