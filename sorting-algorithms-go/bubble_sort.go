package main

import "fmt"

func bubbleSort(input []int) {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			swap(input, j)
		}
	}
}

func swap(slice []int, index int) {
	if slice[index] > slice[index+1] {
		temp := slice[index+1]
		slice[index+1] = slice[index]
		slice[index] = temp
	}
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	bubbleSort(input)
	fmt.Println("AFTER:  ", input)
}
