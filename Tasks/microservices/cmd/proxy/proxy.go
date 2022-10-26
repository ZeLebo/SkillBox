package proxy

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func StartProxy(port1, port2, proxyIP string, signalChan chan os.Signal) {
	proxyIP = "localhost:" + proxyIP

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		client := http.Client{}
		var (
			resp *http.Response
			req  *http.Request
			err  error
		)

		if counter == 0 {
			dest := "http://localhost:" + port1 + r.URL.Path
			req, err = http.NewRequest(r.Method, dest, r.Body)
			if err != nil {
				return
			}
			counter++
		} else {
			dest := "http://localhost:" + port2 + r.URL.Path
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
	})
	fmt.Println("Proxy server started")
	go func() {
		err := http.ListenAndServe(proxyIP, nil)
		if err != nil {
			log.Error(err.Error())
		}
	}()
	<-signalChan
}
