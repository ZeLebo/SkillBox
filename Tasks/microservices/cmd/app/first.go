package main

import (
	log "github.com/sirupsen/logrus"
	c "user/configs"
	s "user/pkg/server"
)

func main() {
	address, err := c.GetIp("first")
	if err != nil {
		log.Fatalln("Cannot parse config")
	}
	server := s.NewServer(address)
	log.Info("The server is up and running at ", server.Addr)

	// signal handler for correct shutdown
	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Info(err.Error())
		}
		done <- true
	}()

	server.WaitShutdown()

	<-done
}
