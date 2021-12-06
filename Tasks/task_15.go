package main

import "fmt"

func evenCounter() (int, int) {
	var (
		arr [10]int
		even int
	)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < 10; i++ {
		if arr[i] % 2 == 0 {
			even++
		}
	}
	return even, 10 - even
}

func reverseArr(arr []int) []int {
	length := len(arr)
	for i := 0; i < length / 2; i++ {
		arr[i], arr[length - i - 1] = arr[length - i - 1], arr[i]
	}
	return arr
}

func secondTask() []int {
	var arr []int
	var tmp int

	_, err := fmt.Scanf("%d", &tmp)
	for err == nil {
		arr = append(arr, tmp)
		_, err = fmt.Scanf("%d", &tmp)
	}

	return reverseArr(arr)
}

func main() {
	// fmt.Println("The first task\nGimme 10 numbers")
	// fmt.Println(evenCounter())
	// fmt.Println("The second task")
	result := secondTask()
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}