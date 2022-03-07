package main

import "fmt"

func ShellSort(data []int) {
	n := len(data)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				data[j], data[j-h] = data[j-h], data[j]
			}
		}
		h /= 3
	}
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	ShellSort(input)
	fmt.Println("AFTER:  ", input)
}
