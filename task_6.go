package main

import (
    "fmt"
)

func taskFirst() {
    println("The first task")
    var a int
    println("Give me the number:")
    fmt.Scanf("%d", &a)
    for i := 1; i < a + 1; i++ {
        fmt.Println(i)
    }
}

func taskSecond() {
    var a, b int
    fmt.Println("Put the numbers:")
    fmt.Scanf("%d %d", &a, &b)
    for a < a + b + 1 {
        a++
    }
    fmt.Println("Here's the sum:\n", a)
}

func taskThird() {
}

func taskForth(){
}

func taskFifth() {
}

func taskSixth() {
}

func main() {
    taskFirst()
    println()
    taskSecond()
    println()
    taskThird()
    println()
    taskForth()
    println()
    taskFifth()
    println()
    taskSixth()
}
