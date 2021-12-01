package main

import (
	"fmt"
	"math/rand"
)

const globalOne = 2
const globalTwo = 5
const globalThree = 10

func isPrime(a int) bool {
	if a % 2 == 0 {
		return true
	} else {
	return false
	}
}

func firstTask() {
	var a int
	fmt.Println("Give me the number")
	size, _ := fmt.Scanf("%d", &a)
	if size == 0 {
		a = rand.Int()
	}
	fmt.Println(isPrime(a))
}

func generatePoints() (ax, ay, bx, by, cx, cy int) {
	return  rand.Intn(1000), rand.Intn(1000),
			rand.Intn(1000), rand.Intn(1000),
			rand.Intn(1000), rand.Intn(1000)
}

func transformPoints(ax, ay, bx, by, cx, cy *int) {
	*ax = 2 * *ax + 10
	*bx = 2 * *bx + 10
	*cx = 2 * *cx + 10
	*ay = -3 * *ay - 5
	*by = -3 * *by - 5
	*cy = -3 * *cy - 5
}

func secondTask() {
	ax, ay, bx, by, cx, cy := generatePoints()
	transformPoints(&ax, &ay, &bx, &by, &cx, &cy)
	fmt.Println(ax, ay)
	fmt.Println(bx, by)
	fmt.Println(cx, cy)
}

func add(a int) (result int) {
	result = a + 25
	return
}

func twice(a int) (result int) {
	result = 2 * a
	return
}

func thirdTask() {
	var a int
	fmt.Println("Give me the number")
	fmt.Scanf("%d", &a)
	a = add(a)
	fmt.Println(twice(a))

}

func one(a int) int {
	return a + globalOne
}

func two(a int) int {
	return a + globalTwo
}

func three(a int) int {
	return a + globalThree
}

func forthTask() {
	var a int
	fmt.Println("Give me the number")
	fmt.Scan(&a)
	fmt.Println(three(two(one(a))))

}

func main() {
	firstTask()
	secondTask()
	thirdTask()
	forthTask()
}
