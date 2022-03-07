package main

import "fmt"

const greeting string = "Hello %s %s, this is your welcome to Go world using text format.\n"

func main() {
	name := getName()
	lastName := "Moeid"
	fmt.Print("Dummy text. \n")
	fmt.Printf("Hello %s, welcome to GO world getting name from standar input.\n", name)
	fmt.Printf(greeting, name, lastName)
}

func getName() string {
	var name string
	name = "Abdul"
	fmt.Print("Type your name: ")
	fmt.Scanf("%s", &name)

	return name
}