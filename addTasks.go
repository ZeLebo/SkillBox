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

/*
1) Написать простой генератор IP адрессов
2) Найти произведение 2 больших натуральных чисел ( N max = 50 )
3) Вывести элементы строки в которой содержаться числа
4) Передача fileName через параметры запуска программы
*/

func IPgenerator() {
    math.seed(1)
}

func main() {
    substring()
    stringSort()
}
