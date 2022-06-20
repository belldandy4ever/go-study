package main

import "fmt"

type bot interface {
	getGreeting() string
}

type enlishBot struct{}
type spanishBot struct{}

func main() {
	eb := enlishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (enlishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
