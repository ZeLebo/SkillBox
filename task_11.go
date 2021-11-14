package main

import (
	"fmt";
	"bufio";
	"os";
	"strconv";
	"strings"
)

func firstTask() {
	fmt.Println("The first task")
	fmt.Println("Put the string to count")

	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')

	if (len(line) == 1 && string(line[0]) == "\n") {
		line = "Go is an Open source programming Language that makes it Easy to build simple, reliable end efficine Software"
	}

	start := true
	cnt := 0

	for i := 0; i < len(line); i++ {
		if string (line[i]) == " " {
			start = true
		} else {
			if start && string (line[i]) >= "A" && string (line[i]) <= "Z" {
				cnt++
			}
			start = false
		}
	}
	fmt.Printf("The sting contains %v words starting with capital letter\n\n", cnt)
}

func trimLastChar ( s string ) string {
    r := []rune(s)
    return string(r[:len(r)-1])
}

func secondTask() {
	var result []string

	fmt.Println("The second task")
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')

	words := strings.Split(line, " ")
	words[len(words) - 1] = trimLastChar(words[len(words) - 1])

	for i := 0; i < len(words); i++ {
		_, err := strconv.Atoi(words[i])
		if err == nil {
			result = append(result, words[i])
		}
	}

	if len(result) != 0 {
		fmt.Println("The string contains numbers:")
		for i := 0; i < len(result); i++ {
			fmt.Printf("%v ", result[i])
		}
		fmt.Println()
	} else {
		fmt.Println("The string has no numbers inside")
	}
}

func main() {
//	firstTask()
	secondTask()
}