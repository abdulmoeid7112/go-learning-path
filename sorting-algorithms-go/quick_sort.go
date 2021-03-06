package main

import "fmt"

func quickSort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}

func recursionSort(data []int, left int, right int) {
	if left < right {
		pivot := partition(data, left, right)
		recursionSort(data, left, pivot-1)
		recursionSort(data, pivot+1, right)
	}
}

func partition(data []int, left int, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}

		for left < right && data[left] <= data[right] {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	quickSort(input)
	fmt.Println("AFTER:  ", input)
}
