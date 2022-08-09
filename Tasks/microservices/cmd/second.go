package main

import (
	log "github.com/sirupsen/logrus"
	s "user/pkg/Server"
)

func main() {
	var address = "0.0.0.0:60495"
	server := s.NewServer(address)
	log.Info("The server is up and running at ", server.Addr, "\n")

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
