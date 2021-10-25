package main

import (
    "fmt"
)


func firstTask() {
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
}

func thirdTask(){
}

func forthTask() {
}

func fifthTask() {
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
