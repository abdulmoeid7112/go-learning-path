package main

import "fmt"

func CountingSort(data []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(data)

	for i := 0; i < length; i++ {
		bucket[data[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			data[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return data
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	CountingSort(input, 9)
	fmt.Println("AFTER:  ", input)
}
