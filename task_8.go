package main

import (
    "fmt"
)


// I don't like this, too many lines
func taskFirst() {
    fmt.Println("The session of the year")
    fmt.Println("Put the month:")
    var month string
    fmt.Scanf("%s", &month)
    fmt.Print("The season is ")
    switch month {
    case "December", "January", "February":
        fmt.Println("Winter")

    case "March", "April", "May":
        fmt.Println("Spring")

    case "June", "July", "August":
        fmt.Println("Summer")

    case "September", "October", "November":
        fmt.Println("Fall")

    default:
        fmt.Println("something I dont know")
    }
}

func taskSecond() {
    fmt.Println("Days of week")
    fmt.Println("Put the weekday: mon, tue, wed, thu, fri")
    day := ""
    fmt.Scanf("%s", &day)
    switch day {
    case "mon":
        fmt.Println("Monday")
        fallthrough
    case "tue":
        fmt.Println("Tuesday")
        fallthrough
    case "wed":
        fmt.Println("Wednesday")
        fallthrough
    case "thu":
        fmt.Println("Thursday")
        fallthrough
    case "fri":
        fmt.Println("Friday")
        break
    }
}

func lemonadeChange(bills []int) bool {
    five, ten := 0, 0
    for i := 0; i < len(bills); i++ {
        switch bills[i]{
        case 5:
            five++
        case 10:
            if five > 0 {
                five--
                ten++
            } else {
                return false
            }
        case 20:
            if five > 1 && ten > 0 {
                five--
                ten--
            } else if five > 2 {
                five -= 3
            } else {
                return false
            }
        default:
            return false
        }
    }
    return true
}

func taskThird() {
    fmt.Println("The change calculation")
    fmt.Println("Put the bill:")
    var cash int
    var bills []int

    _, result := fmt.Scanf("%d", &cash)

    for result == nil {
        bills = append(bills, cash)
        _, result = fmt.Scanf("%d", &cash)
    }

    fmt.Println(lemonadeChange(bills))
}

func main() {
    taskFirst()
    println()
    taskSecond()
    println()
    taskThird()
}
