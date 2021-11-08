package main

import (
	"fmt";
	"math"
)

func taskFirst() {
	fmt.Println("The first task")
	var num8 uint8 = 0
	var num16 uint16 = 0
	var num32 uint32 = 0
	var cnt8, cnt16 int = 0, 0

	for num32 = 0; num32 < math.MaxUint32; num32++ {
		if num8 == math.MaxUint8 {
			cnt8++
			num8 = 0
		}
		if num16 == math.MaxUint16 {
			cnt16++
			num16 = 0
		}
		num8++
		num16++
	}

	fmt.Println("The amount of unit8 overflow", cnt8)
	fmt.Println("The amount of unit16 overflow", cnt16)
}

func taskSecond() {
	fmt.Println("The second task")
	var num1, num2 int16
	fmt.Scan(&num1, &num2)

	var result int32 = int32(num1) * int32(num2)

	if result > -1 {
		if result < 256 {
			fmt.Println("uint8")	
		} else if (result < 65536) {
			fmt.Println("uint16")
		} else {
			fmt.Println("uint32")
		}
	} else {
		if result > -129 && result < 128 {
			fmt.Println("int8")
		} else if result > -32769 && result < 32768 {
			fmt.Println("int16")
		} else {
			fmt.Println("int32")
		}
	}
}

func main() {
	taskFirst()
	taskSecond()
}