package main

import (
	"fmt";
    "math"
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

func max(a, b float32) float32 {
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

    if ( a == b && b == c) {
        println("They are all the same")
    } else {
        fmt.Print(max(max(a, b), c))
        print(" ")
        if (a == max(max(a,b), c)) {
            a = -1.0
        } else if ( b == max(max(a, b), c)) {
            b = -1.0
        } else {
            c = -1.0
        }
        fmt.Print(max(max(a,b), c))
        print(" \n")
    }
}

func sixthTask() {
    var a, b, c float64
    println("The sixth task")
    println("Put the a, b and c of the equation")
    fmt.Scanf("%f %f %f", &a, &b, &c)
    d := (math.Pow(b, 2) - 4 * a * c)
    if d < 0 {
        println("No solutions")
        return
    } else if ( d == 1 ) {
        var solution float64
        solution = ((-b + math.Sqrt(d)) / 2 * a)
        fmt.Println(solution)
        return
    } else {
        var solution1, solution2 float64
        solution1 = ((-b + math.Sqrt(d)) / 2 * a)
        if solution1 == -0 {
            solution1 = 0
        }
        solution2 = ((-b - math.Sqrt(d)) / 2 * a)
        if solution2 == -0 {
            solution2 = 0
        }
        if (solution1 != solution2) {
            fmt.Println(solution1, solution2)
        } else {
            fmt.Println(solution1)
        }
        return
    }
}

func seventhTask() {
    var num [4]int
    numberString, answer := "0000", "обычный"
    println("The seventh task")
    println("Put the number:")
    fmt.Scanf("%s", &numberString)
    for i := 0; i < 4; i++ {
        num[i] = int(numberString[i]) - 48
    }
    if ( num[0] == num[3] && num[1] == num[2] ) {
        answer = "зеркальный"
        fmt.Println(numberString, "->", answer, "билет")
        return
    } else if ( num[0] + num[1] == num[2] + num[3]) {
        answer = "счастлийвый"
    }
    fmt.Println(numberString, "->", answer, "билет")
}

func eightTask() {
    var answer string = "y"
    num := 8

    println("The eight task")
    for i := 4; i >= 1; i /= 2 {
        fmt.Println("Is your number >", num,"[y/n] ?")
        fmt.Scanf("%s", &answer)

        if (answer == "y") {
            num += i
        } else {
            num -= i
        }
    }
    fmt.Println("Is your number >", num, "[y/n] ?")
    fmt.Scanf("%s", &answer)
    if (answer == "y") {
        num++
    }
    fmt.Println("Your number is", num)
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
