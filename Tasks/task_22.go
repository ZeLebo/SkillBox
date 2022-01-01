package main

import (
	"fmt"
	"math/rand"
)

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[pivot], arr[right]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	qSort(arr[:left])
	qSort(arr[left+1:])

	return arr
}

func binarySearch(n, left, right int, arr []int) int {
	mid := (left + right) / 2

	if left == mid {
		if arr[right] != n && arr[left] != n {
			return -1
		} else if arr[right] == n {
			return right
		} else if arr[left] == n {
			return left
		}
	}

	if arr[mid] == n {
		for arr[mid-1] == n {
			mid--
		}
		return mid
	} else if arr[mid] < n {
		return binarySearch(n, mid, right, arr)
	} else if arr[mid] > n {
		return binarySearch(n, left, mid, arr)
	}
	return -1
}

func main() {
	fmt.Println("Generator of random numbers")
	var randomArr [10]int
	for i := 0; i < 10; i++ {
		randomArr[i] = rand.Int() % 100
	}
	fmt.Println(randomArr)
	var a int
	fmt.Scan(&a)
	count := -1
	for i := 0; i < 10; i++ {
		if randomArr[i] == a {
			count = i
			break
		}
	}
	if count == -1 {
		fmt.Println("There no such number")
	} else {
		fmt.Printf("After %v in the array %v numbers\n", a, 9-count)
	}

	fmt.Println("Quick sort and binary search")
	arr := []int{1, 2, 3, 4, 4, 4, 7, 8, 9, 10, 11, 12}
	fmt.Println("Give me the number to find in array")
	fmt.Scan(&a)
	arr = qSort(arr)
	fmt.Println(binarySearch(a, 0, len(arr)-1, arr))
}
