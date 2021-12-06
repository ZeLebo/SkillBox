package main

import (
	"fmt"; 
	"math"
)

func firstTask() {
	fmt.Println("The first task")
	var (
		x, prevResult, epsilon, step float64
		result = float64(1)
		fact = float64 (1)
		n int
	)
	fmt.Println("Put x")
	fmt.Scan(&x)
	fmt.Println("Put the needed precision")
	fmt.Scan(&n)
	epsilon = 1.0 / math.Pow(10, float64 (n))

	for step = 1; math.Abs(result - prevResult) > epsilon; step++ {
		prevResult = result
		fact *= step
		result += math.Pow(x, step) / fact
	}

	fmt.Println(result)
}

func secondTask() {
	var (
		years int
		amount, percent, forBank float64

	)
	fmt.Println("The second task")
	fmt.Println("Write the amount of money, percent and amount of years")
	fmt.Scanf("%f %f %d", &amount, &percent, &years)

	for i := 0; i < years * 12; i++ {
		amount *= 1 + percent / 100
		forBank += amount - math.Floor(amount * 100) / 100
		amount = math.Floor(amount * 100) / 100
	}
	fmt.Println("The money you get:", amount)
	fmt.Println("Money the bank gets:", forBank)
}


func main() {
	firstTask()
	secondTask()
}