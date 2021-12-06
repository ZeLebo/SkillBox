package main

import (
	"fmt"
)

func taskFirst() {
    println("The first task")
    var a int
    println("Give me the number:")
    fmt.Scan(&a)
    for i := 0; i < a + 1; i++ {
        fmt.Print(i, " ")
    }
}

func taskSecond() {
    var a, b int
    fmt.Println("The second task")
    fmt.Println("Put the numbers:")
    fmt.Scanf("%d %d", &a, &b)
    for i := 0 ; i < b; i++ {
        a++
    }
    fmt.Println("Here's the sum:", a)
}

func taskThird() {
    var price, discount float32
    fmt.Println("The third task")
    fmt.Println("Put the price and discount:")
    fmt.Scanf("%f %f", &price, &discount)
    discount /= 100
    if discount == 0 && price == 0 {
        fmt.Println("You don't have to pay")
        return
    }
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
    a, b, c := 0, 0, 0
    for true {
        if a < 10 {
            a++
        }
        if b < 100 {
            b++
        }
        if c < 1000 {
            c++
        }
        if a == 10 && b == 100 && c == 1000 {
            break
        }
    }
}

type Bucket struct {
    full bool
    cnt, cap int32
}

func summary (bucket [3]Bucket) int {
    result := 0
    for i := 0; i < 3; i++ {
        result += int (bucket[i].cap)
    }
    return result
}

func taskFifth() {
    fmt.Println("The fifth task")
    var bucket [3]Bucket
    for i := 0; i < 3; i++ {
        bucket[i].full = false
        bucket[i].cnt = 0
        fmt.Print("Give the capacity of the ", i + 1 ," bucket: ")
        fmt.Scanf("%d", &bucket[i].cap)

        if ( bucket[i].cap == 0 ) {
            fmt.Println("What a strange bucket… UwY")
            fmt.Println("I don't play with you anymore…")
            break
        } else {
            for bucket[i].cnt < bucket[i].cap {
            bucket[i].cnt++
            }    
        }
    }
    if summary(bucket) != 0 {
        for i := 0; i < 3; i++ {
            fmt.Print(bucket[i].cnt, " apples in ", i + 1," bucket of ", bucket[i].cap, " available", "\n")
        }
    }
}

func sum(arr [3]int) int {
    result := 0
    for _, v := range arr {
        result += v
    }
    return result
}

func taskSixth() {
    peopleOnFloors := [3]int{3, 3, 3}
    peopleInLift := 0

    fmt.Println("People on the floors:", peopleOnFloors)

    for sum(peopleOnFloors) != 0 {
        peopleInLift = 0
        fmt.Println("Going upwards and downwards and take 2 people")

        for i := 2; i > -1; i-- {
            if (peopleInLift == 2) {
                break
            }

            if (peopleOnFloors[i] > 1 && peopleInLift  == 0) {
                peopleOnFloors[i] -= 2
                peopleInLift += 2
                break

            } else if (peopleOnFloors[i] > 1 ){
                peopleOnFloors[i] -= 1
                peopleInLift++
                break 

            } else if (peopleOnFloors[i] == 1 || peopleInLift == 1) {
                peopleOnFloors[i]--
                peopleInLift++
            } else {
                continue
            }
        }
    fmt.Println("People on the floors:", peopleOnFloors)
    }
}

func main() {
    taskFirst()
    println()
    taskSecond()
    println()
    taskThird()
    taskForth()
    println()
    taskFifth()
    println()
    taskSixth()
}
