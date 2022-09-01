package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"user/configs"
)

var counter int

// todo split
// redirects the output
func handle(w http.ResponseWriter, r *http.Request) {

	client := http.Client{}
	var (
		resp *http.Response
		req  *http.Request
	)

	addr1, err := configs.GetIp("first")
	if err != nil {
		log.Fatalln("cannot parse config")
	}
	addr2, err := configs.GetIp("second")
	if err != nil {
		log.Fatalln("cannot parse config")
	}

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
