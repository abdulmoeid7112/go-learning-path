package main

import "fmt"

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	InsertionSort(input)
	fmt.Println("AFTER:  ", input)
}
