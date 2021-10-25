package main

import (
    "fmt"
)


func firstTask() {
    print("The first task\n")
    var x, y float32
    fmt.Println("Put the x and y coordinates:")
    fmt.Scanf("%f %f", &x, &y)
    if ( x == 0 || y == 0) {
        print("I don't know how to determine it\n")
    } else if ( x > 0 && y > 0 ) {
        print("1 part\n")
    } else if (x > 0 && y < 0) {
        print("4 part\n")
    } else if (x < 0 && y > 0) {
        print("2 part\n")
    } else {
        print("3 part\n")
    }
 }

func secondTask() {
    var a, b, c float32
    print("The second task\nPut 3 numbers here:\n")
    fmt.Scanf("%f %f %f", &a, &b, &c)
    if (a > 0 || b > 0 || c > 0) {
        print("Something is bigger that zero\n")
    } else {
        print("The numbers are all not positive\n")
    }
}

func thirdTask() {
    var a, b, c float32
    print("The third task\nPut 3 numbers here:\n")
    fmt.Scanf("%f %f %f", &a, &b, &c)
    if ( a == b || a == c || c == b ) {
        print("Some of the numbers are the same\n")
    } else {
        print("All the numbers are different\n")
    }
}

func forthTask() {
    var toPay int
    var coin [3]int
    println("The forth task")
    fmt.Println("Put the value and denomination of 3 coins, you wanna pay with:")
    fmt.Scanf("%d %d %d %d", &toPay, &coin[0], &coin[1], &coin[2])

    for  i := 0; i < 3; i++ {
        for toPay > coin[i] {
            toPay %= coin[i]
        }
    }
    if toPay == 0 {
        print("You can pay without change\n")
    } else {
        print("You have to take change\n")
    }
}

func max(float32 a, float32 b) {
    if a > b {
        return a
    }
    return b
}

func fifthTask() {
    var a, b, c float32
    println("The fifth task")
    println("Put the interest rates ( without % symbol )")
    fmt.Scanf("%f %f %f", &a, &b, &c)
    print(max(max(a, b), c))
}

func sixthTask() {
}

func seventhTask(){
}

func eightTask() {
}


func main() {
    firstTask()
    fmt.Println()
    secondTask()
    fmt.Println()
    thirdTask()
    fmt.Println()
    forthTask()
    fmt.Println()
    fifthTask()
    fmt.Println()
    sixthTask()
    fmt.Println()
    seventhTask()
    fmt.Println()
    eightTask()
}
