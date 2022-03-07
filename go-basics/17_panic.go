package main

import (
	"errors"
	"fmt"
)

func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("the value A is a string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("the value B is a string")
	}

	return a.(int) + b.(int), nil
}

func main() {
	number, err := Sum(50, 50)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sun: %d\n", number)

	_, err = Sum("50", 50)
	if err != nil {
		panic(err)
	}
}
