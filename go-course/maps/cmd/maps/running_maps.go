package maps

import "fmt"

func Run() {
	var cars map[string]string

	cars["BMW"] = "320i"
	cars["BMW"] = "x1"
	cars["BMW"] = "x6"
	cars["Mercedes"] = "c180"

	for key, value := range cars {
		fmt.Printf("Marca: %v - Modelo: %v", key, value)
	}
}