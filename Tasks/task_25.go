package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func errorHandler(err error) {
    if err != nil {
        panic(err)
    }
}

func noErrorAtoi(num string) int {
    ans, _ := strconv.Atoi(num)
    return ans
}

func getArrayUser() []int {
    fmt.Println("Give me the numbers")
    arrString, err := bufio.NewReader(os.Stdin).ReadString('\n'); errorHandler(err)
    var result []int
    for _, num := range strings.Fields(arrString) {
        result = append(result, noErrorAtoi(num))
    }
    return result
}

func main() {
    res := getArrayUser()
    fmt.Println(res)
}
