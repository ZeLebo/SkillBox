package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func foo(ch chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println("New number " + string(<-ch))
	}
}

func main() {
	fmt.Println("main() started")
	long := make(chan int, 3)
	long <- 3
	long <- 4
	long <- 5
	go foo(long)

	c := make(chan string)

	go greet(c)
	c <- "John"

	go greet(c)
	c <- "ZhoRa"


	fmt.Println("main() stopped")
}