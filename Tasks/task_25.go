package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func getArrayUser() []int {
    fmt.Println("Give me the numbers")
    arrString, err := bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
        panic(err)
    }

    var result []int
    hateTheseErrors := strings.Fields(arrString)
    for _, num := range hateTheseErrors {
        convNum, _ := strconv.Atoi(num)
        result = append(result, convNum)
    }

    return result
}

func main() {
    res := getArrayUser()
    fmt.Println(res)
}
