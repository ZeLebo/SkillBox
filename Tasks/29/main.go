package main

import (
    "fmt"
    "strconv"
)

func main() {
    ch := make(chan int, 1)
    sqr := make(chan int, 1)
    mul := make(chan int, 1)
    var (
        word string
        num int
    )
    _, err := fmt.Scanf("%s", &word)
    for err == nil && word != "stop" && word != "стоп" {
        num, err = strconv.Atoi(word)
        if err == nil {
            ch <- num
            fmt.Println("Input:", <-ch)
            sqr <- num * num
            fmt.Println("Square:", <-sqr)
            mul <- num * num * 2
            fmt.Println("Multiply:", <- mul)
        }
        _, err = fmt.Scanf("%s", &word)
    }
}
