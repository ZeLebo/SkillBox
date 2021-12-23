package main

import "fmt"

func deferMaking(something func (int, int) int) {
	defer something(2, 2)
}

func calculations() {
	var (
		s, z float32
		x int16
		y uint8
	)
	fmt.Println("Put x, y and z:")
	fmt.Scan(&x, &y, &z)
	fmt.Println(3 / z)
	s = float32(x << 1) + float32(y * y) - 3.0/z
	fmt.Println("The result is:", s)
}

func main() {
	fmt.Println("The first task")
	calculations()

	defer func(a, b int) {
		fmt.Println(a, "+", b, "=",a + b)
	}(1, 2)

	foo := func(a, b int) int {fmt.Println(a, "-", b, "=", a - b); return a - b}
	defer foo(1, 2)

	deferMaking(func(a, b int) int {fmt.Println(a, "*", b, "=", a * b); return a * b})
}
