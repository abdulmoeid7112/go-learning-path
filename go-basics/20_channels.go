package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		ch <- "Hi Channel"
	}()

	fmt.Println(<-ch)

	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
	}()

	for n := range ch1 {
		fmt.Printf("%d\n", n)
	}

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	//ch2 <- 3
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)

	c := make(chan string)
	// Starting a concurrent goroutine
	go greet(c)
	// Sending values to the channel c
	c <- "World"
	// Closing channel
	close(c)
}

func greet(c chan string) {
	name := <-c // receiving value from channel
	fmt.Println("Hello", name)
}
