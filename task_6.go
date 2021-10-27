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
