package main

import "fmt"

func tripleDet(matrix [3][3]int) int {
	result := 0
	for i := 0; i < 3; i++ {
		result += matrix[0][i] * matrix[1][(i+1)%3] * matrix[2][(i+2)%3]
		result -= matrix[0][i] * matrix[1][(i+2)%3] * matrix[2][(i+1)%3]
	}

	return result
}

func multiply(matrix1 [3][5]int, matrix2 [4][5]int) [3][4]int {
	var result [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 5; k++ {
				result[i][j] += matrix1[i][k] * matrix2[j][k]
			}
		}
	}

	return result
}

func main() {
	fmt.Println("The first task")
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		_, err := fmt.Scanf("%d %d %d", &matrix[i][0], &matrix[i][1], &matrix[i][2])
		if err != nil {
			print("Something went wrong")
		}
	}

	fmt.Println(tripleDet(matrix))

	fmt.Println("The second task")

	var (
		matrix1 = [3][5]int{
			{1, 2, 3, 12, 23},
			{4, 5, 6, 12, 12},
			{7, 7, 8, 12, 12},
		}
		matrix2 = [4][5]int{
			{10, 20, 30, 32, 12},
			{23, 34, 12, 23, 12},
			{14, 44, 23, 32, 28},
			{14, 44, 23, 32, 228},
		}
	)

	result := multiply(matrix1, matrix2)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

}
