package main
import (
    "fmt";
    "strings";
    "sort"
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

func main() {
    substring()
    stringSort()
}
