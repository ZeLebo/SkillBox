package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func firstTask() {
	fmt.Println("Write the input into the file")
	var message string

	input := bufio.NewReader(os.Stdin)
	message, err := input.ReadString('\n')
	errorHandler(err)
	dt := time.Now()

	if message != "exit\n" {
		f, _ := os.Create("test.txt")

		for i := 1; message != "exit\n"; i++ {
			dt = time.Now()
			_, err = f.WriteString(strconv.Itoa(i) + ". " + dt.Format("2006-01-02 15:04:05") + " " + message)
			if err != nil {
				return
			}
			message, _ = input.ReadString('\n')
		}
	} else {
		return
	}
}
func secondTask() {
	fmt.Println("Read the data from file")

	file, fileOpenError := os.Open("test.txt")
	if fileOpenError != nil {
		fmt.Println("The file wasn't found")
		return
	}

	input := bufio.NewReader(file)
	fileInfo, err := os.Stat("test.txt")
	errorHandler(err)

	if fileInfo.Size() == 0 {
		fmt.Println("The file is empty")
	} else {
		fileStrings, fileReadError := input.ReadString('\n')
		for fileReadError == nil {
			fmt.Print(fileStrings)
			fileStrings, fileReadError = input.ReadString('\n')
		}
	}
}

func thirdTask() {
	fmt.Println("Change the file permission")
	f, err := os.Create("permission.txt")
	errorHandler(err)

	stats, _ := os.Stat("permission.txt")
	fmt.Printf("Permision before: %s\n", stats.Mode())
	_, _ = f.WriteString("Hello worlds")
	err = os.Chmod("permission.txt", 777)
	errorHandler(err)

	stats, _ = os.Stat("permission.txt")
	err = os.Remove("permission.txt")
	errorHandler(err)

	fmt.Printf("Permision after: %s\n", stats.Mode())
}

func forthTask() {
	fmt.Println("The forth task")
	fmt.Println("Put the data here")
	answer := ""

	os.Create("anotherTest.txt")
	var stringToWrite string
	for i := 1; true; i++ {
		answer, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		if answer == "exit\n" {
			ioutil.WriteFile("anotherTest.txt", []byte(stringToWrite), 666)
			break
		}
		stringToWrite += strconv.Itoa(i) + ". " + time.Now().Format("2006-01-02 15:04:05") + " " + answer
	}

	fmt.Println("Reading data from file")
	file, err := ioutil.ReadFile("anotherTest.txt")
	errorHandler(err)
	fmt.Printf("%s", file)

	os.Remove("anotherTest.txt")
}

func fifthTask() {
	fmt.Println("Round brackets combinations")
	var (
		amount = 3
		answer = ""
	)
	fmt.Println("Put the number of pairs of brackets")
	_, err := fmt.Scanf("%d", &amount)
	if err != nil {
		return
	}
	str := make([]rune, amount*2)
	result := generateParens(amount, amount, str, 0)

	fmt.Printf("The amount of possible combinations is %d"+
		"\nDo you wanna see the result itself?[yes/no]\n", len(result))
	if _, err = fmt.Scan(&answer); err == nil && answerParser(answer) {
		fmt.Println("Okay, here it is:")
		for i, s := range result {
			fmt.Printf("%d. %s\n", i+1, s)
		}
	} else if !answerParser(answer) {
		fmt.Println("Don't wanna waste your time")
	} else {
		fmt.Printf("Problem is: %s", err)
	}
}

func answerParser(answer string) bool {
	if strings.ToLower(strings.Trim(answer, "!?.,;: \n")) == "yes" {
		return true
	} else {
		return false
	}
}

func generateParens(leftRem int, rightRem int, str []rune, count int) []string {
	var result []string
	if leftRem < 0 || rightRem < leftRem {
	} else if leftRem == 0 && rightRem == 0 {
		result = append(result, string(str))
	} else {
		if leftRem > 0 {
			str[count] = '('
			result = append(result, generateParens(leftRem-1, rightRem, str, count+1)...)
		}
		if rightRem > 0 {
			str[count] = ')'
			result = append(result, generateParens(leftRem, rightRem-1, str, count+1)...)
		}
	}
	return result
}

func main() {
	firstTask()
	secondTask()
	thirdTask()
	forthTask()
	fifthTask()
}
