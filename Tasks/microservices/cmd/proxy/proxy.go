package main

import (
	env "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var counter int

func init() {
	err := env.Load("../../.env")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

// todo split
// redirects the output
func handle(w http.ResponseWriter, r *http.Request) {
	addr1 := os.Getenv("FIRST_IP")
	addr2 := os.Getenv("SECOND_IP")
	client := http.Client{}
	var (
		resp *http.Response
		req  *http.Request
		err  error
	)

	if counter == 0 {
		dest := "http://" + addr1 + r.URL.Path
		req, err = http.NewRequest(r.Method, dest, r.Body)
		if err != nil {
			return
		}
		counter++
	} else {
		dest := "http://" + addr2 + r.URL.Path
		req, err = http.NewRequest(r.Method, dest, r.Body)
		if err != nil {
			return
		}
		counter--
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Error("Error: ", err.Error())
	}

	content, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(content)
	if err != nil {
		log.Error(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}

func main() {
	proxyIP := os.Getenv("PROXY_IP")
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(proxyIP, nil))
}
