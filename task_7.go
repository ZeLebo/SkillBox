package main

import (
    "fmt"
)

func taskFirst() {
    cnt, start, end := 0, 100000, 999999
    fmt.Println("The amount of mirror tickets")
    fmt.Println("from 100000 to 999999")
    for i := start; i <= end; i++ {
        if (i % 1000 == (i - i % 1000) / 1000) {
            cnt++
        }
    }
    fmt.Println(cnt)
}

func taskSecond() {
    var height, width int
    var str1, str2 string
    fmt.Println("Chess board")
    fmt.Println("Put the width:")
    fmt.Scan(&width)
    if (width <= 0) {
        fmt.Println("Starnge field")
        return
    } else {
        for i := 0; i < width; i++ {
            if (i % 2 == 0) {
                str1 += " "
                str2 += "*"
            } else {
                str1 += "*"
                str2 += " "
            }
        }
    }

    fmt.Println("Put the height:")
    fmt.Scan(&height)

    if (height < 1) {
        fmt.Println("Starnge field")
        return
    } else {
        for i := 0; i < height; i++ {
            if (i % 2 == 0) {
                fmt.Println(str1)
            } else {
                fmt.Println(str2)
            }
        }
    }
}

func taskThird() {
    fmt.Println("Chrismas tree printing")
    fmt.Println("Put the height:")
    var height int
    var str string
    fmt.Scan(&height)
    if (height < 1) {
        fmt.Println("Starnge tree")
        return
    } else {
        for i := 0; i < height; i++{
            str = ""
            for j := 0; j < height * 2; j++ {
                if j >= height - i && j <= height + i {
                    str += "*"
                } else {
                    str += " "
                }
            }
            fmt.Println(str)
        }
    }
}

func summary(num int) int {
    result := 0
    for num > 0 {
        result += num % 10
        num /= 10
    }
    return result
}

func taskForth() {
    fmt.Println("Lucky tickets")
    cnt, prev, start, end := 0, 100001, 100000, 999999
    fmt.Print("The amount of tickets needed to buy: ")
    for i := start; i <= end; i++ {
        if (summary(i % 1000)) == summary((i - i % 1000) / 1000) {
            if ( cnt < i - prev) {
                cnt = i - prev
            }
            prev = i
        }
    }
    fmt.Println(cnt)
}

func main() {
    taskFirst()
    println()
    taskSecond()
    println()
    taskThird()
    println()
    taskForth()
}
