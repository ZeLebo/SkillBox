package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user/cmd/handlers"
	"user/internal/database"
	"user/internal/domain"
	"user/internal/service"
)

func main() {
	logger := log.Default()
	cfg := domain.GetDatabaseConfig()

	client, err := database.NewClient(cfg, logger)
	if err != nil {
		log.Fatalf("error initializing database: %s", err.Error())
	}

	handlerService := service.NewService(client, logger)

	mainRouter := mux.NewRouter()
	handler := handlers.NewRequestHandler(handlerService, logger)

	handler.Routes(mainRouter)

	address, err := domain.GetIp("first")
	if err != nil {
		log.Fatalln("Cannot parse config")
	}
	go func() {
		err := http.ListenAndServe(address, mainRouter)
		if err != nil {
			logger.Fatal("Can't start server" + err.Error())
		}
	}()
	logger.Printf("Server started")
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-signalChan
	logger.Println("Shutting down server...")

	defer os.Exit(0)
}
