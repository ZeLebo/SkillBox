package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user/internal/database"
	"user/internal/domain"
	"user/internal/handlers"
	"user/internal/service"
)

func StartServer(port string, signalChan chan os.Signal) {
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

	address := "localhost:" + port
	go func() {
		err := http.ListenAndServe(address, mainRouter)
		if err != nil {
			logger.Fatal("Can't start server" + err.Error())
		}
	}()
	logger.Printf("Server started")

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
