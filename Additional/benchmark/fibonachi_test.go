package benchmark

import (
		"fmt";
		"testing";
		"math/rand"
)

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

func Benchmark_recursion(b *testing.B) {
	n := rand.Intn(40)
	fmt.Println(recursion(n))

	// fmt.Println(recursion(30))
}

func Benchmark_addition(b *testing.B) {
	n := rand.Intn(40)
	fmt.Println(addition(n))

	// fmt.Println(addition(30))
}
