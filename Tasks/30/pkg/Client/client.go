package user

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	my "user/pkg/Functions"
)

func connectServer(mode, host string) net.Conn {
	d, err := net.Dial(mode, host)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func ListenFromServer() {
	d := connectServer("tcp", "localhost:8080")

	for {
		// send info from stdin to server
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		my.ErrorHandler(err)
		_, err = d.Write([]byte(line))
		my.ErrorHandler(err)

		// get the info from the server
		text, err := bufio.NewReader(d).ReadString('\n')
		my.ErrorHandler(err)

		// print the data from server
		fmt.Println(line[:len(line)-1], "->", string(text[:len(text)-1]))
	}
}
