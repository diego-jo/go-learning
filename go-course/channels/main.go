package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("started")
	links := []string{
		"https://goolge.com",
		"https://youtube.com",
		"https://twitter.com",
		"https://nasa.gov",
		"https://www.mercadolibre.com.mx",
	}

	for _, link := range links {
		go checklink(link)
	}

	// dúvidas:
	// posso colocar outras instruções após lançar uma go routine
	// e por último espera o retorno dos channels???

	// uma go routine pode chamar outras? de forma encadeada?
}

func checklink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("%s is down\n", link)
		return
	}

	fmt.Printf("%s is up\n", link)
}
