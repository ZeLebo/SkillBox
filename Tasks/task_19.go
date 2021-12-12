package main

import "fmt"

func merge(arr1 []int, arr2 []int) []int {
	var result []int
	i, j := 0, 0

	for k := 0; k < len(arr1)+len(arr2); k++ {
		if i == len(arr1) || j == len(arr2) {
			continue
		} else {
			if arr1[i] < arr2[j] {
				result = append(result, arr1[i])
				i++
			} else {
				result = append(result, arr2[j])
				j++
			}
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr1) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func main() {
	fmt.Println("Merging")
	arr1 := []int{1, 2, 5, 6}
	arr2 := []int{3, 4, 5, 7, 10}
	fmt.Println(merge(arr1, arr2))

	fmt.Println("Bubble sort")

	arr := []int{5, 192, 34, 234, 54, 23, 12, -34, 7329, 1983, 2734, 1, 3, 4}

	for range arr {
		for j := range arr {
			if j == len(arr)-1 {
				continue
			}
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
