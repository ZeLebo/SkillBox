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
    for i := 0 ; i < b; i++ {
        a++
    }
    fmt.Println("Here's the sum:", a)
}

func taskThird() {
    var price, discount float32
    fmt.Println("Put the price and discount:")
    fmt.Scanf("%f %f%", &price, &discount)
    discount /= 100
    if ( discount > 0.3 ) {
        discount = 0.3
    } else if (discount < 0.0 ){
        discount = 0.0
    }
    if (discount * price) > 2000 {
        price -= 2000
    } else {
        price *= ( 1.0 - discount )
    }
    fmt.Println("Here's how much you have to pay:", price)
}


func taskForth() {

    fmt.println()
    fmt.Println("Here's the number:")
    fmt.Println("Here's the number of dfk")
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
