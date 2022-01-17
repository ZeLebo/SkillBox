package main

import (
    "fmt"
    "os"
)

/*
can use flags for input -> go run main.go --str "some" --substr "s"
OR use without flags -> go run main.go "Hello" "lo"
*/
func parseInput() (str, substr string) {
    if len(os.Args) < 2 {return}
    if os.Args[1] == "--str" && os.Args[3] == "--substr" {
        for i, v := range os.Args {
            if i == len(os.Args) {break}

            if v == "--str" && os.Args[i + 1] != "--substr" {
                str = os.Args[i + 1]
            } else if v == "--substr" {
                substr = os.Args[i + 1]
            }
        }
    } else {
        str, substr = os.Args[1], os.Args[2]
    }
    return
}

func contains(strOld, substrOld string) bool {
    str, substr := []rune(strOld), []rune(substrOld)

    if len(substr) == 0 && len(str) == 0 { return true }
    if substr == nil || len(substr) == 0 || str == nil || len(str) == 0 { return false }

    for i := 0; i < len(str); i++ {
        if str[i] == substr[0] {
            for j := 0; j < len(substr); j++ {
                if i > len(str) - 1 {
                    return false
                }
                if str[i] != substr[j] {
                    return false
                }
                i++
            }
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(contains(parseInput()))
}
