package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMap())
	fmt.Println(getMap2())
	fmt.Println(getMap3("Berlin"))
}

func getMap() map[string]int {
	mapTest := make(map[string]int)
	mapTest["key1"] = 1
	mapTest["key2"] = 100

	return mapTest
}

func getMap2() map[string]int {
	mapTest := map[string]int{
		"Tokyo":     18,
		"Lisbon":    24,
		"Professor": 30,
	}

	delete(mapTest, "Juan") // not exist

	return mapTest
}

func getMap3(name string) int {
	mapTest := map[string]int{
		"Helsinki":  18,
		"Berlin":    24,
		"Stockholm": 30,
	}

	delete(mapTest, "Stockholm") // exist

	return mapTest[name]
}
