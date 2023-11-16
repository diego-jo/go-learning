package main

import "fmt"

type bot interface {
	// Apenas por criar uma interface aqui e add uma func com essa assinatura
	// todos os tipos struct que tem a mesma assinatura, passam a ter acesso
	// a func printGreeting como sendo "herança" da interface
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type greet []string //não apenas tipos struct entram nessa regra

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
