package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("I've made it")
	}()

	a := func(a int) int {
		return a * a
	}(3)
	fmt.Println(a)
}
