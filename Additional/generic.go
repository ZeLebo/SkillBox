package main

import "fmt"

func main() {
    fmt.Println(min(1,2))
    fmt.Println(min(0.1 , 0.2))

}

func min[T interface{ int | float32 | float64 }](a T, b T) T {
    if a < b {
        return a
    } else {
        return b
    }
}

