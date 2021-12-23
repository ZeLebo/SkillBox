package main
import "fmt"

func mySwap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func additionHere(a, b int) int {
	var result int
	if a > b {
		mySwap(&a, &b)
	}
	for i := a; i < b + 1; i++ {
		if i % 2 == 0 {
			result += i
		}
	}
	return result
}

func main() {
	var a, b int
	fmt.Println("The first task")
	fmt.Println("Put here 2 numbers:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(additionHere(a, b))

	fmt.Println("The second task")
	fmt.Println("Before swap:", a, b)
	mySwap(&a, &b)
	fmt.Println("After swap:", a, b)
}
