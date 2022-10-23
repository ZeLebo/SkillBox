package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"user/internal/domain"
)

var counter int

func handle(w http.ResponseWriter, r *http.Request) {

	client := http.Client{}
	var (
		resp *http.Response
		req  *http.Request
	)

	addr1, err := domain.GetIp("first")
	if err != nil {
		log.Fatalln("cannot parse config")
	}
	addr2, err := domain.GetIp("second")
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

	//goland:noinspection GoDeprecation
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
	proxyIP, err := domain.GetIp("proxy")
	if err != nil {
		log.Fatalln("Cannot start the proxy server")
	}
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(proxyIP, nil))
}
