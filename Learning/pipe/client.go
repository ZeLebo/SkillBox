package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"fmt"
)

func trimLastChar(str string) string {
	r := []rune(str)
	return string(r[:len(r) - 1])
}

func main() {
	d, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if _, err := d.Write([]byte(line)); err != nil {
			log.Fatal(err)
		}

		text, err := bufio.NewReader(d).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(trimLastChar(line), "->", string(text))
	}
}
