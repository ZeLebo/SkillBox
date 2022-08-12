package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("../../logs/logs.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		file, err = os.Create("../../logs/logs.log")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(file)
	//log.SetOutput(file)
}
