package main

import "fmt"

//if any other type w getgreeting that rets a string,
//then that type is an honorary member of type bot
//and can call printgreeting of type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

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
	//Very custom logic for generating enlish greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola Amigo!"
}
