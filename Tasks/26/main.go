package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "io/ioutil"
)

func readFromFile(filename string) string {
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    return string(b)
}

func writeToFile(filename string, strs []string) {
    file, _ := os.Create(filename)
    defer file.Close()
    ioutil.WriteFile(filename, []byte(strings.Join(strs, "")), 777)
}

func main() {
    var str []string
    if len(os.Args) == 2 {
        fmt.Print(readFromFile(os.Args[1]))
    } else if len(os.Args) == 3 {
        str = append(str, readFromFile(os.Args[1]), readFromFile(os.Args[2]))
        fmt.Print(strings.Join(str, "\n"))
    } else if len(os.Args) == 4 {
        str = append(str, readFromFile(os.Args[1]), readFromFile(os.Args[2]))
        writeToFile(os.Args[3], str)
    }
}
