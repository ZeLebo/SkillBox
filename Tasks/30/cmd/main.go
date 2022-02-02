/*
This file need to be started first, than it will do all the work
*/
package main

import (
	client "user/pkg/client"
	server "user/pkg/server"
)

func main() {
	go server.ListenFromClient()
	client.ListenFromServer()
}
