package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func square() {
    var num int
    for {
        _, err := fmt.Scanf("%d", &num)
        if err == nil {
            fmt.Println("square:", num * num)
        }
    }
}


func main() {
    go square()

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    
    signal.Notify(sigs, syscall.SIGINT)

    go func() {
        fmt.Println(<- sigs)
        done <- true
    } ()

    <- done
    fmt.Println("Exiting...")
}
