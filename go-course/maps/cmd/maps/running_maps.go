package main

import "fmt"

func run() {
	// cars := map[string]string{
	// 	"bmw": "320i",
	// 	"audi": "TT",
	// }

	// var cars map[string]string
	// cars = make(map[string]string)
	cars := make(map[string]string)

	cars["BMW"] = "320i"
	cars["BMW"] = "x6"
	cars["Mercedes"] = "c180"

	for key, value := range cars {
		fmt.Printf("Marca: %v - Modelo: %v\n", key, value)
	}

	// elem, ok := cars["BNW"]

	// fmt.Println(elem, ok)
}
