package main

import (
	env "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	s "user/pkg/server"
)

func init() {
	err := env.Load("../../.env")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	address := os.Getenv("FIRST_IP")
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
