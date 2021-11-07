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
    case "December":
        fmt.Println("Winter")
    case "January":
        fmt.Println("Winter")
    case "February":
        fmt.Println("Winter")

    case "March":
        fmt.Println("Spring")
    case "April":
        fmt.Println("Spring")
    case "May":
        fmt.Println("Spring")

    case "June":
        fmt.Println("Summer")
    case "July":
        fmt.Println("Summer")
    case "August":
        fmt.Println("Summer")

    case "September":
        fmt.Println("Fall")
    case "October":
        fmt.Println("Fall")
    case "November":
        fmt.Println("Fall")

    default:
        fmt.Println("something I dont know")
    }
}

func taskSecond() {
    fmt.Println("Days of week")
    fmt.Println("Put the weekday: mon, tue, wed, thu, fri")
    day := ""
    fmt.Scanf("%d", &day)
    days := [5]string {"Monday", "Thuesday", "Wednesday", "Thuesday", "Friday"}
    fmt.Println(days)
    switch day {
    case "mon":
        fmt.Println(day)
        for i := 0; i < 5; i++ {
            fmt.Println(days[i])
        }
    case "tue":
        fmt.Println(day)
        for i := 1; i < 5; i++ {
            fmt.Println(days[i])
        }
    case "wed":
        fmt.Println(day)
        for i := 2; i < 5; i++ {
            fmt.Println(days[i])
        }
    case "thu":
        fmt.Println(day)
        for i := 3; i < 5; i++ {
            fmt.Println(days[i])
        }
    case "fri":
        fmt.Println(day)
        for i := 4; i < 5; i++ {
            fmt.Println(days[i])
        }
    default:
        fmt.Println("It's not the day of the suggested above")
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
