package main

import (
	"fmt"

	"github.com/abdulmoeid7112/go-learning-path/go-basics/mypackage"
)

func main() {
	firstName := mypackage.GetName()
	fmt.Println("Your name is: ", firstName)
}
