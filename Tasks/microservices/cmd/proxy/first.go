package main

import (
	"errors"
	env "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	s "user/pkg/server"
)

func init() {
	err := env.Load("../../.env")
	if err != nil {
		// if not found, need to find the file
		filename, err := func(name string) (string, error) {
			var result string
			root := "../../"
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Error(err)
					return err
				}
				if !info.IsDir() && filepath.Ext(path) == name {
					result = path
					return nil
				}
				return errors.New("not found")
			})
			return result, err
		}(".env")

		if err != nil {
			log.Error(err)
			os.Exit(1)
		} else {
			err = env.Load(filename)
			if err != nil {
				log.Error(err)
			}
		}

		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
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
