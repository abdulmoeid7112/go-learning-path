package main

import "fmt"

func main() {
	a, b, c := getVars()
	fmt.Println(a, b, c)
}

func getVars() (int, int, int) {
	return 1, 2, 3
}
