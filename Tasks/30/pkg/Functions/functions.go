package user

import (
	"log"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TrimLastChar(str string) string {
	tmp := []rune(str)
	return string(tmp[:len(tmp)-1])
}

func RemoveIndex(arr []interface{}, index int) []interface{} {
	return append(arr[:index], arr[index+1:]...)
}

func Index(arr []interface{}, something interface{}) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == something {
			return i
		}
	}
	return -1
}

func Remove(arr []interface{}, something interface{}) []interface{} {
	return RemoveIndex(arr, Index(arr, something))
}