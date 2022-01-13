package main

import (
		"fmt";
		"testing"
)

func Benchmark_recursion()

func recursion (n int) int {
	result := 0
	if n < 1 {
		result = 0
	} else if n == 1 {
		result = 0
	} else if n == 2 {
		result = 1
	} else {
		result = recursion(n-1) + recursion(n-2)
	}
	return result
}

func addition(n int) int {
	result := 0
	if n == 1 {
		result = 0
	} else if n == 2 {
		result = 1
	} else {
		num1, num2 := 0, 1
		for i := 2; i < n; i++ {
			result = num2 + num1
			num1 = num2
			num2 = result
		}
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(addition(n))
	fmt.Println(recursion(n))
}
