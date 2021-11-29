package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "sort"
    "strconv"
    "strings"
    "time"
)

func substring(){
    var sentence, word string
    fmt.Println("Give me the sentence:")
    fmt.Scanf("%s", &sentence)
    fmt.Println("Give me the word to check:")
    fmt.Scanf("%s", &word)
    fmt.Println(strings.Contains(sentence, word))
}

func stringSort(){
    words := []string{"", "", ""}
    fmt.Println("Give me the worlds:")
    fmt.Scanf("%s %s %s", &words[0], &words[1], &words[2])
    sort.Strings(words)
    for i := 0; i < 3; i++ {
        fmt.Println(words[i])
    }
}

func IPgenerator() {
    fmt.Println("IP generator")
    var ip string
    rand.Seed(time.Now().UnixMilli())
    for i := 0; i < 4; i++ {
        ip += strconv.Itoa(rand.Int() % 999)
        if i != 3 {
            ip += "."
        }
    }
    fmt.Println("Your IP is", ip)
}

func runesToString(runes [100]rune) (outString string) {
    var index int
    for i := 0; i < 100; i++ {
        if runes[i] != '0' - 48 {
            index = i
            break
        }
    }

    for i := index; i < 100; i++ {
        outString += string(runes[i] + 48)
    }
    return outString
}

func longMultiply() {
    fmt.Println("Give me 2 numbers to multiply")
    var num1, num2 string
    fmt.Scanf("%s %s", &num1, &num2)
    var (
        first, second [50]rune
        result [100]rune
    )
    for i := 0; i < len(num1); i++ {
        first[49 - i] = rune(num1[len(num1) - i - 1]) - 48
    }
    for i := 0; i < len(num2); i++ {
        second[49 - i] = rune(num2[len(num2) - i - 1]) - 48
    }
    for i := 0; i < 25; i++ {
        first[i], first[49 - i] = first[49 - i], first[i]
        second[i], second[49 - i] = second[49 - i], second[i]
    }

    for i := 0; i < 50 ; i++ {
        for j := 0; j < 50; j++ {
            result[j + i] += first[j] * second[i]
        }
    }

    for i := 0; i < 99; i++ {
        result[i + 1] += result[i] / 10
        result[i] %= 10
    }
    for i := 0; i < 50; i++ {
        result[i], result[99 - i] = result[99 - i], result[i]
    }

    fmt.Println(runesToString(result))
}

func checkString(str string, subs ...string) bool {
    for _, sub := range subs {
        if strings.Contains(str, sub) {
            return true
        }
    }
    return false
}

func numberContatining() {
    fmt.Println("Numbers finding")
    line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    stringToFind := strings.Split(line, " ")
    for i := range(stringToFind) {
        if checkString(stringToFind[i], "0", "1", "2", "3", "4", "5", "6", "7", "8", "9") {
            fmt.Println(stringToFind[i])
        }
    }
}

func fileNameChanging(fileName string) {
    file, _ := os.Create(fileName)
    defer file.Close()
    stringToWrite := "Hello world!"
    ioutil.WriteFile(fileName, []byte(stringToWrite), 755)
}

func main() {
    //substring()
    //stringSort()
    //IPgenerator()
    longMultiply()
    //numberContatining()
    //fileNameChanging(os.Args[1])
}
