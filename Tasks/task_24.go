package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
    "strconv"
	"math/rand"
)

func getArrayUser() []int {
    fmt.Println("Give me the numbers")
    arrString, err := bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
        panic(err)
    }

    var result []int
    for _, num := range strings.Fields(arrString) {
        convNum, _ := strconv.Atoi(num)
        result = append(result, convNum)
    }

    return result
}

const ARR_SIZE = 10
func selectionSort(a [ARR_SIZE]int) [ARR_SIZE]int {
    for j := 2; j < ARR_SIZE; j++ {
        key := a[j]; i := j - 1
        for i > 0 && a[i] > key {
            a[i + 1] = a[i]; i--
        }
        a[i + 1] = key
    }
    return a
}

func main() {
	var randomArr [ARR_SIZE]int
	fmt.Println("Give me the numbers")
	for i := 0; i < ARR_SIZE; i++ {
		randomArr[i] = rand.Int() % 100
	}
    fmt.Println("Selection sort:", selectionSort(randomArr))

    fmt.Println("Wrong bubble (right stone) sort:",
        func (a... int) [] int {
            for i := 0; i < len(a); i++ {
                for j := 0; j < len(a); j++ {
                    if a[i] > a[j] {
                        a[i], a[j] = a[j], a[i]
                    }
                }
            }
            return a
        } (getArrayUser()...))
}
