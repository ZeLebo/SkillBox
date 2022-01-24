package main

import (
    "fmt"
    "strconv"
    "sync"
)

func inputPrint(num int, wg *sync.WaitGroup) int {
    defer wg.Done()
    fmt.Println("Input:", num)
    return num * num
}

func square(num int, wg *sync.WaitGroup) int {
    defer wg.Done()
    fmt.Println("Square: ", num)
    return num * 2
}

func multiply(num int) {
    fmt.Println("Multiply:", num)
}

func main() {
    var (
        word string
        num int
        err error
        wg sync.WaitGroup
    )

    _, err = fmt.Scanf("%s", &word)
    for err == nil && word != "stop" && word != "стоп" {
        num, err = strconv.Atoi(word)
        if err == nil {
            wg.Add(2)
            go func () {
                fs := inputPrint(num, &wg)
                sc := square(fs, &wg)
                multiply(sc)
            } ()
            wg.Wait()
        }
        _, err = fmt.Scanf("%s", &word)
    }
}
