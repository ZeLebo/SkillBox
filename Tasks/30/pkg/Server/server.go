package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	my "user/pkg/Functions"
)

func startServer(mode, host string) net.Conn {
	listener, err := net.Listen(mode, host)
	my.ErrorHandler(err)
	fmt.Println("The server is listening on port: 8080")

	con, err := listener.Accept()
	my.ErrorHandler(err)
	return con
}

func ListenFromClient() {
	con := startServer("tcp", "localhost:8080")

	for {
		// get the info from the client
		line, err := bufio.NewReader(con).ReadString('\n')
		my.ErrorHandler(err)
		// work with the info
		upper := strings.ToUpper(string(line))
		// send the info to the client
		_, err = con.Write([]byte(upper))
		my.ErrorHandler(err)
	}
}
