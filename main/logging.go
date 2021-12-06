package main

import (
    "fmt"
    "time"

    log "github.com/sirupsen/logrus"
)

func runAndWait() int {
    time.Sleep(time.Millisecond * 1)
    return 10
}

func init() {
    log.SetFormatter(&log.JSONFormatter{})

    log.SetLevel(log.WarnLevel)
}

func main() {
    for i := 0; i < 100; i++ {

        a := runAndWait()
        log.Info("runAndWait finished...")
        log.Infof("result = %d", a)

        if i > 90 {
            log.WithFields(log.Fields{
                "i": i,

            }).Warn("closer to the edge")
        }

        if i == 99 {
            log.Fatal("reached 100")
        }
    }

    fmt.Println("done")
}
