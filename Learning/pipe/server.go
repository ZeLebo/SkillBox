package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("The server is running")

	con, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		line, err  := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		upper := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upper)); err != nil {
			log.Fatalln(err)
		}
	}

}
