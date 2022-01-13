package main

import (
    "fmt";
    "strings"
)

const ARR_SIZE = 10
const SENT_SIZE = 4
const CHARS_SIZE = 5

func firstTask() {
    testArr := [ARR_SIZE]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    even, odd := evenOdd(testArr)
    if len(even) == 0 && len(odd) == 0 {
        panic("Something went wrong…")
    }
    fmt.Println("The even array:", even)
    fmt.Println("The odd array:", odd)
}

func isEven(a int) bool {
    if a % 2 == 0 {
        return true
    }
    return false
}

func evenOdd(arr [ARR_SIZE]int) ([] int, []int) {
    var even, odd []int

    for i := range(arr) {
        if isEven(arr[i]) {
            even = append(even, arr[i])
        } else {
            odd = append(odd, arr[i])
        }
    }
    return even, odd
}

func secondTask() {
    sentences := [SENT_SIZE]string {
        "Hello world", "Hello ZhoRa",
        "Turn the lighst off", "Congrats Skillbox"}
    chars := [CHARS_SIZE]rune {'H', 'O', 'L', 'W', 'S'}

    result := parseTest(sentences, chars)
    for i := range result {
        for j := range result[i] {
            if result[i][j] != -1 {
                fmt.Println(string(chars[j]), "found in sentence", i + 1 ,"at index:", result[i][j])
            }
        }
        println()
    }
}

func getLastWordIndex(sentence string) int {
    if len(sentence) == 0 {
        return 0
    }
    for i := len(sentence) - 1; i > -1; i-- {
        if sentence[i] == ' ' {
            return i + 1
        }
    }
    return 0
}

func findIdx(sentence string, char rune) int {
    id := getLastWordIndex(sentence)
    for ; id < len(sentence); id++ {
        if string(sentence[id]) == string(char) || string(sentence[id]) == strings.ToLower(string(char)) {
            return id
        }
    }
    return -1
}

func checkLetter(a byte, b rune) bool {
    if string(a) == string(b) || string(a) == strings.ToLower(string(b)) {
        return true
    }
    return false
}

func parseTest(sentences [SENT_SIZE]string, chars [CHARS_SIZE]rune) ([SENT_SIZE][CHARS_SIZE]int) {
    if len(sentences) == 0 || len(chars) == 0 {
        panic("Wrong data set")
    }

    var result [SENT_SIZE][CHARS_SIZE]int

    for i := 0; i < SENT_SIZE; i++ {
        for j := 0; j < CHARS_SIZE; j++ {
            result[i][j] = -1
        }
    }

    for i := range sentences {
        for j := range chars {
            result[i][j] = findIdx(sentences[i], chars[j])
        }
    }
    return result
}

func main() {
    fmt.Println("The first task:"); firstTask()
    fmt.Println("\nThe second task:"); secondTask()
}
