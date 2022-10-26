package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"user/cmd/app"
	"user/cmd/proxy"
)

func main() {
	fmt.Println("Put one number for port to start server")
	fmt.Println("And put 3 numbers for ports (server1 IP, server2 IP, proxy IP) to start proxy")
	fmt.Println()
	if len(os.Args) == 1 {
		log.Fatal("Need to provide the port(s) to start")
	}
	var flagServer = true

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		os.Kill)

	if len(os.Args) > 3 {
		fmt.Println("This is proxy")
		flagServer = false
	} else {
		fmt.Println("This is server")
	}

	if flagServer {
		// server spawn
		port := os.Args[1]
		app.StartServer(port, signalChan)
	} else {
		// proxy with servers spawn
		handlersChan := make(chan os.Signal, 3)
		port1, port2, proxyIP := os.Args[1], os.Args[2], os.Args[3]
		fmt.Println("Starting two servers")

		go app.StartServer(port1, handlersChan)
		go app.StartServer(port2, handlersChan)
		proxy.StartProxy(port1, port2, proxyIP, handlersChan)
		// catch signal, if caught -> push 2 more signals to channel
		<-signalChan
		for i := 0; i < 3; i++ {
			handlersChan <- os.Interrupt
		}
	}
}
