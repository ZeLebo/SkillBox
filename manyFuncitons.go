package main
import "fmt"

func sum(nums ... int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func multiply(nums ...int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}

// It's working, but the functions are passed by value
// the functions are not passed as links
// the result of function are passed to that function
// and I cannot understand how to pass them into this function and execute inside the function
func multiFuncResult(f ...int) {
	for i := len(f) - 1; i > -1; i-- {
		fmt.Println("The result of", i + 1, "function is", f[i])
	}
}

// Make it execute after linking or compiling( forgot this moment )
// Here I can pass only exact amount of arguments
func manyFunc(a, b int, f ... func(... int) int) {
	for i := len(f) - 1; i > -1; i-- {
		fmt.Println(f[i](a, b))
	}
}

// There no way to pass 2 varang, so this "costyl" is making it possible
// Making Map, but with my hands…
func funcLink(f []func(... int) int, nums ... int) {
	for _, exec := range f {
		fmt.Println(exec(nums...))
	}
}
func functions(f ... func(... int) int) []func(... int) int {
	return f
}

func main() {
	funcLink(functions(sum, multiply), 1, 2 ,3, 4, 5, 6, 7)
	//multiFuncResult(sum(1, 2, 4), multiply(1, 2, 4))
	//manyFunc(2, 2, sum, multiply)
}
