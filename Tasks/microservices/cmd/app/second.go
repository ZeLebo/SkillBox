package main

import (
	log "github.com/sirupsen/logrus"
	"user/configs"
	s "user/pkg/server"
)

func main() {
	address, err := configs.GetIp("second")
	if err != nil {
		log.Fatalln("cannot parse config")
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
