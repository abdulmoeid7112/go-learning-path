package main

import "fmt"

func insertionSort(data []int) {
	for i := 0; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
}

func BucketSort(data []int, bucketSize int) []int {
	var max, min int
	for _, n := range data {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range data {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {
	input := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
		fmt.Scan(&input[i])
	}
	fmt.Println("BEFORE: ", input)

	fmt.Println("AFTER:  ", BucketSort(input, 2))
}
