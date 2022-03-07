package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomQuickSort(list []int, start, end int) {
	if end-start > 1 {
		mid := randomPartition(list, start, end)
		RandomQuickSort(list, start, mid)
		RandomQuickSort(list, mid+1, end)
	}
}

func randomPartition(list []int, begin, end int) int {
	rand.Seed(time.Now().UnixNano())
	i := begin + rand.Intn(end-begin)
	list[i], list[begin] = list[begin], list[i]
	return partition2(list, begin, end)
}

func partition2(list []int, begin, end int) (i int) {
	cValue := list[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if list[j] < cValue {
			i++
			list[j], list[i] = list[i], list[j]
		}
	}
	list[i], list[begin] = list[begin], list[i]
	return i
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	RandomQuickSort(input, 0, 10)
	fmt.Println("AFTER:  ", input)
}
