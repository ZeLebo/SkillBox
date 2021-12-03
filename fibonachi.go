package main

import "fmt"

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

func main() {
	var n, result int
	fmt.Scan(&n)
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
	fmt.Println(result)
	fmt.Println(recursion(n))
}
