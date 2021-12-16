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

func main() {
	fmt.Println("The first task")
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d %d %d", &matrix[i][0], &matrix[i][1], &matrix[i][2])
	}

	fmt.Println(tripleDet(matrix))

	fmt.Println("The second task")
	var (
		matrix1 [5][5]int
		matrix2 [5][5]int
		result  [5][5]int
	)

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			if i != 3 {
				fmt.Scanf("%d", &matrix1[i][j])
			}
			fmt.Scanf("%d", &matrix2[i][j])
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

}
