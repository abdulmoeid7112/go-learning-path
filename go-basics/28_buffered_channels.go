package main

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg3.Done()
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go foo(fooVal, i)
	}
	wg3.Wait()
	close(fooVal)
	for item := range fooVal {
		fmt.Println(item)
	}
}
