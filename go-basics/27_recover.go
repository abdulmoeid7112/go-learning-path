package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
	wg2.Done()
}

func say2(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear... a 2")
		}
	}
}

func main() {
	wg2.Add(1)
	go say2("Hey")
	wg2.Add(1)
	go say2("there")
	wg2.Wait()
}
