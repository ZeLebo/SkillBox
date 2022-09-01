package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

type HTTPErrorHandler struct {
	ErrorCode   int
	Description string
}

func HTTPErrorHandle(w http.ResponseWriter, err HTTPErrorHandler) {
	w.WriteHeader(err.ErrorCode)
	// If the logger is on server, then log it
	if err.ErrorCode == http.StatusInternalServerError {
		log.Error(err.Description)
	}
	_, err1 := w.Write([]byte(err.Description))
	if err1 != nil {
		return
	}
	return
}

func RootDir() string {
	ex, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot parse the working directory")
	}
	exPath := filepath.Base(ex)
	switch exPath {
	case "microservices":
		return "logs"
	case "cmd":
		return filepath.Join("..", "logs")
	case "app":
		return filepath.Join("..", "..", "logs")
	default:
		return ""
	}
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	logPath := filepath.Join(RootDir(), "logs.log")
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("no such file but why")
		file, err = os.Create(logPath)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(file)
	//log.SetOutput(file)
}
