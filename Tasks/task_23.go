package main

import "fmt"

const ARR_SIZE = 10

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

func firstTask() {
    testArr := [ARR_SIZE]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    even, odd := evenOdd(testArr)
    if even == nil || odd == nil {
        fmt.Println("Something went wrong…")
    }
}

const SENT_SIZE = 4
const CHARS_SIZE = 5

func secondTask() {
    sentences := [SENT_SIZE]string {
        "Hello world",
        "Hello Skillbox",
        "Привет Мир",
        "Привет Skillbox"}

    chars := [CHARS_SIZE]rune {
        'H', 'E', 'L', 'П', 'М'}
    for _, word := range(sentences) {
        getLastWord(word)
    }
    fmt.Println(parseTest(sentences, chars))
}

func getLastWord(sentence string) (int, string) {
    if len(sentence) == 0 {
        panic("The word is too small")
    }

    var word []rune

    for i := len(sentence) - 1 ; i > -1 && sentence[i] != ' '; i-- {
        word = append(word, rune(sentence[i]))
    }
    for i := 0; i < len(word) / 2; i++ {
        word[i], word[len(word) - i - 1] = word[len(word) - i - 1], word[i]
    }
    return len(sentence) - len(word), string(word)
}

func getLastWordIndex(sentence string) int {
    if len(sentence) == 0 {
        return 0
    }
    for i := len(sentence) - 1; i > -1; i-- {
        if sentence[i] == ' ' {
            return i
        }
    }
    return 0
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

    for i := 0; i < SENT_SIZE; i++ {
        for index := getLastWordIndex(sentences[i]); index < len(sentences[i]); index++ {
            print(string(sentences[i][index]))
            for j := range chars {
                if string(sentences[i][index]) == string(chars[j]) {
                    result[i][j] = index
                    break
                }
            }
        }
    }
    return result
}


func main() {
    firstTask(); secondTask();
}
